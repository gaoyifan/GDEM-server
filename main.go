package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/golang-lru"
	"gopkg.in/redis.v3"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var (
	imageCache *lru.Cache
	imageLock  map[Point]*sync.Mutex
	mapCache   *redis.Client
)

const (
	prefix         string = "map"
	imageCacheLen  int    = 20
	minZoom        uint   = 9
	imageLength    int    = 360
	degreeToSecond        = 3600
	listenAddress  string = ":8000"
)

type Point struct {
	lon, lat               float64
	serialLon, serialLat   int //latitude and longitude for map file
	subSerialX, subSerialY int //0.1*0.1 degree image location in 1*1 degree area
	pixelX, pixelY         int //pixel in one image
}

func (p *Point) genMapInfo() {
	p.serialLon = int(p.lon)
	p.serialLat = int(p.lat)
	if p.lon <= 0 {
		p.serialLon = -p.serialLon
		p.serialLon++
	}
	if p.lat < 0 {
		p.serialLat = -p.serialLat
		p.serialLat++
	}
	pixelInOneDegreeX := ((int)((p.lon-math.Floor(p.lon))*degreeToSecond) + degreeToSecond) % degreeToSecond
	pixelInOneDegreeY := (degreeToSecond - (int)((p.lat-math.Floor(p.lat))*degreeToSecond)) % degreeToSecond
	p.subSerialX = pixelInOneDegreeX / imageLength
	p.subSerialY = pixelInOneDegreeY / imageLength
	p.pixelX = pixelInOneDegreeX % imageLength
	p.pixelY = pixelInOneDegreeY % imageLength
}

func newPoint(lon, lat float64) *Point {
	p := new(Point)
	p.lon = lon
	p.lat = lat
	p.genMapInfo()
	return p
}

func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	imageCache, err = lru.New(imageCacheLen)
	if err != nil {
		fmt.Println(err)
	}
	imageLock = make(map[Point]*sync.Mutex)

	mapCache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	mapCache.ConfigSet("save", "60 10")
}

func XYToLonLat(xtile, ytile int, zoom uint) *Point {
	var lon_deg, lat_deg float64
	b := (1 << zoom)
	n := float64(b)
	lon_deg = (float64)(xtile*360)/n - 180
	lat_deg = math.Atan(math.Sinh(math.Pi*(1-2*(float64)(ytile)/n))) * 180 / math.Pi
	return newPoint(lon_deg, lat_deg)
}

func lnglatToXY(p *Point, zoom int) (int, int) {
	lat_rad := p.lat * math.Pi / 180
	n := 1 << (uint)(zoom)
	xtile := (int)(math.Floor((p.lon + 180) / 360 * (float64)(n)))
	ytile := (int)((1 - math.Log(math.Tan(lat_rad)+1/math.Cos(lat_rad))/math.Pi) * (float64)(n) / 2)
	return xtile, ytile
}

func getImageFileName(p *Point) string {
	var (
		lat_dir, lon_dir string
	)
	if p.serialLat > 0 {
		lat_dir = "N"
	} else {
		lat_dir = "S"
	}
	if p.serialLon >= 0 {
		lon_dir = "E"
	} else {
		lon_dir = "W"
	}
	return fmt.Sprintf("%s/%s%02d/%s%03d/%1d%1d", prefix, lat_dir, p.serialLat, lon_dir, p.serialLon, p.subSerialX, p.subSerialY)
}

var funLock sync.Mutex

func getImage(p *Point) *[imageLength][imageLength]int16 {
	var imgLock *sync.Mutex
	var imgLockExist bool

	funLock.Lock()
	imgLock, imgLockExist = imageLock[*p]
	if !imgLockExist {
		imgLock = new(sync.Mutex)
		imageLock[*p] = imgLock
	}
	funLock.Unlock()

	imgLock.Lock()
	defer imgLock.Unlock()

	var img *[imageLength][imageLength]int16
	imgInterface, cached := imageCache.Get(p)
	if cached {
		img = imgInterface.(*[imageLength][imageLength]int16)
	} else {
		file, err := os.Open(getImageFileName(p))
		if err != nil {
			return nil
		}
		defer file.Close()

		fileByte, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fileBuffer := bytes.NewReader(fileByte)

		img = new([imageLength][imageLength]int16)
		for i := 0; i < imageLength; i++ {
			for j := 0; j < imageLength; j++ {
				err := binary.Read(fileBuffer, binary.LittleEndian, &img[i][j])
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
		}
		imageCache.Add(p, img)
	}
	return img
}

func getMap(i0, j0 int, zoom uint, size_index int) []byte {
	var (
		latSpan, lonSpan float64
		w                bytes.Buffer
		img *[imageLength][imageLength]int16
	)
	pStart := XYToLonLat(i0, j0, zoom)
	pEnd := XYToLonLat(i0+1, j0+1, zoom)
	fmt.Printf("%f\t%f\t%f\t%f\n", pStart.lat, pStart.lon, pEnd.lat, pEnd.lon)
	size := 1 << (uint)(size_index)
	latSpan = (pEnd.lat - pStart.lat) / (float64)(size)
	lonSpan = (pEnd.lon - pStart.lon) / (float64)(size)

	if zoom < minZoom {
		for i := 0; i <= size; i++ {
			for j := 0; j <= size; j++ {
				binary.Write(&w, binary.LittleEndian, int16(0))
			}
		}
	} else {
		for i := 0; i <= size; i++ {
			var p *Point
			//p.lat = pEnd.lat - (float64)(i)*latSpan
			for j := 0; j <= size; j++ {
				//p.lon = pStart.lon + (float64)(j)*lonSpan
				p=newPoint(pStart.lon + (float64)(j)*lonSpan,pEnd.lat - (float64)(i)*latSpan)
				p.genMapInfo()
				img = getImage(p)
				if img == nil {
					binary.Write(&w, binary.LittleEndian, int16(0))
					continue
				}
				binary.Write(&w, binary.LittleEndian, img[p.pixelX][p.pixelY])
			}
		}
	}
	return w.Bytes()
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	var (
		i0, j0, size_index, zoomi int
		zoom                      uint
		err                       error
		response                  []byte
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
	fmt.Printf("%s\n", r.URL)
	responseCache := mapCache.Get(r.URL.String())

	if responseCache.Err() == redis.Nil { // not cached
		response = getMap(i0, j0, zoom, size_index)
		mapCache.Set(r.URL.String(), response, 0)
	} else if responseCache.Err() != nil { // redis error or not connected
		response = getMap(i0, j0, zoom, size_index)
		fmt.Println(responseCache.Err())
	} else { // cached
		response, err = responseCache.Bytes()
	}

	// write response
	if len(response) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/{i:[0-9]+}/{j:[0-9]+}/{zoom:[0-9]+}/{size:[0-9]+}", mapHandler).Methods("GET")
	http.Handle("/", route)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		fmt.Println(err)
	}
}
