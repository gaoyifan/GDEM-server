package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePng, err := os.Open(os.Args[1])
	re, _ := regexp.Compile("ASTGTM2_(...)(....)_dem")
	submatch := re.FindSubmatch([]byte(os.Args[1]))
	lat := (string)(submatch[1])
	lon := (string)(submatch[2])
	fmt.Println(lat,lon)

	if err != nil {
		fmt.Println(err)
	}
	defer filePng.Close()
	img, err := png.Decode(filePng)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var w bytes.Buffer

	for a := 0; a < 10; a++ { // cut one image to 100 dat file (10 * 10 mesh)
		for b := 0; b < 10; b++ {
			fileDat,err:=os.Create(os.Args[2]+"/"+lat+lon+strconv.Itoa(a)+strconv.Itoa(b))
			if err !=nil{
				fmt.Println(err)
				os.Exit(1)
			}
			for i := a * 360; i < (a+1)*360; i++ {
				for j := b * 360; j < (b+1)*360; j++ {
					gray, _, _, _ := img.At(i, j).RGBA()
					binary.Write(&w, binary.LittleEndian, int16(gray))
				}
			}
			w.WriteTo(fileDat)
			fileDat.Close()
		}
	}
}
