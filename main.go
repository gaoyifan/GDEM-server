package main

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/golang-lru"
	"image"
	"image/png"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

var (
	prefix        string = "map/ASTGTM2_"
	suffix        string = "_dem.png"
	imageCache    *lru.Cache
	imageCacheLen int = 20
)

type Point struct {
	lon, lat float64
}

type PointInt struct {
	lon, lat int
}

func (c Point) toInt() PointInt {
	var r PointInt
	r.lat = (int)(c.lat)
	r.lon = (int)(c.lon)
	return r
}
func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	imageCache, err = lru.New(imageCacheLen)
	if err != nil {
		fmt.Println(err)
	}
}

func XYToLonLat(xtile, ytile int, zoom uint) Point {
	var lon_deg, lat_deg float64
	var rt Point
	b := (1 << zoom)
	n := float64(b)
	lon_deg = (float64)(xtile*360)/n - 180
	lat_deg = math.Atan(math.Sinh(math.Pi*(1-2*(float64)(ytile)/n))) * 180 / math.Pi
	rt.lon = lon_deg
	rt.lat = lat_deg
	return rt
}

func lnglatToXY(longitude, latitude float64, zoom int) (int, int) {
	lat_rad := latitude * math.Pi / 180
	n := 1 << (uint)(zoom)
	xtile := (int)(math.Floor((longitude + 180) / 360 * (float64)(n)))
	ytile := (int)((1 - math.Log(math.Tan(lat_rad)+1/math.Cos(lat_rad))/math.Pi) * (float64)(n) / 2)
	return xtile, ytile
}

func getImageFileName(p PointInt) string {
	var (
		lat_dir, lon_dir string
	)
	if p.lat >= 0 {
		lat_dir = "N"
	} else {
		lat_dir = "S"
	}
	if p.lon >= 0 {
		lon_dir = "E"
	} else {
		lon_dir = "W"
	}
	return fmt.Sprintf("%s%s%02d%s%03d%s", prefix, lat_dir, p.lat, lon_dir, p.lon, suffix)
}

func getImage(p Point) image.Image {
	pInt := p.toInt()
	img, cached := imageCache.Get(pInt)
	if !cached {
		file, err := os.Open(getImageFileName(pInt))
		if err != nil {
			return nil
		}
		defer file.Close()
		img, err = png.Decode(file)
		if err != nil {
			return nil
		}
		imageCache.Add(pInt, img)
	}
	return img.(image.Image)
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	var (
		i0, j0, size_index, zoomi int
		zoom                      uint
		latSpan, lonSpan          float64
		err                       error
	)
	para := mux.Vars(r)
	if i0, err = strconv.Atoi(para["i"]); err != nil {
		return
	}
	if j0, err = strconv.Atoi(para["j"]); err != nil {
		return
	}
	if zoomi, err = strconv.Atoi(para["zoom"]); err != nil {
		return
	}
	zoom = uint(zoomi)
	if size_index, err = strconv.Atoi(para["size"]); err != nil {
		return
	}
	pStart := XYToLonLat(i0, j0, zoom)
	pEnd := XYToLonLat(i0+1, j0+1, zoom)
	fmt.Printf("%f\t%f\t%f\t%f\t", pStart.lat, pStart.lon, pEnd.lat, pEnd.lon)
	fmt.Println(r.URL)
	size := 1 << (uint)(size_index)
	latSpan = (pEnd.lat - pStart.lat) / (float64)(size)
	lonSpan = (pEnd.lon - pStart.lon) / (float64)(size)
	for i := 0; i < size; i++ {
		var p Point
		p.lat = pEnd.lat - (float64)(i)*latSpan
		for j := 0; j < size; j++ {
			p.lon = pStart.lon + (float64)(j)*lonSpan
			img := getImage(p)
			if img == nil {
				binary.Write(w, binary.LittleEndian, int16(0))
				continue
			}
			x := (int)((p.lon - math.Floor(p.lon)) * (float64)(img.Bounds().Max.X))
			y := img.Bounds().Max.Y - (int)((p.lat-math.Floor(p.lat))*(float64)(img.Bounds().Max.Y))
			gray, _, _, _ := img.At(x, y).RGBA()
			binary.Write(w, binary.LittleEndian, int16(gray))
		}
	}
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/{i:[0-9]+}/{j:[0-9]+}/{zoom:[0-9]+}/{size:[0-9]+}", mapHandler).Methods("GET")
	http.Handle("/", route)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
