package main

import (
	"fmt"
	"image"
	"strconv"

	_ "image/jpeg"
	"image/png"
	"os"

	_ "golang.org/x/image/tiff"
)

func main() {

	args := os.Args

	if len(args) <= 6 || len(args) > 7 {
		fmt.Println("Use by doing:\nimageCutter inputImageWithExtension x1 y1 x2 y2 outputImageWithoutExtension")
		return
	}

	imgRead, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer imgRead.Close()

	// image.Decode requires that you import the right image package. We've
	// imported "image/png", so Decode will work for png files. If we needed to
	// decode jpeg files then we would need to import "image/jpeg".
	//
	// Ignored return value is image format name.
	img, _, err := image.Decode(imgRead)
	if err != nil {
		panic(err)
	}

	out, _ := os.Create(args[6] + ".png")
	defer out.Close()

	x1, _ := strconv.ParseInt(args[2], 10, 64)
	y1, _ := strconv.ParseInt(args[3], 10, 64)
	x2, _ := strconv.ParseInt(args[4], 10, 64)
	y2, _ := strconv.ParseInt(args[5], 10, 64)

	// type assertion
	editable := img.(*image.RGBA)

	actualSubImage := editable.SubImage(image.Rect(int(x1), int(y1), int(x2), int(y2)))

	// in case you want to set specific pixels it needs to be converted back to the struct
	// or to an interface like draw.Image
	// pixelSettable := actualSubImage.(*image.RGBA)
	// pixelSettable.Set(104, 105, color.RGBA{255, 0, 0, 255})

	png.Encode(out, actualSubImage)

}
