package main

import (
	"image"
	"os"
	"image/png"
)

func createPNG(img *image.RGBA) {
	file, err := os.Create("test.png")
	check(err)
	defer file.Close()

	err = png.Encode(file, img)
	check(err)
}

func main() {
	//img := createIMG1()
	dx := 256
	dy := 256
	cv := newCanvas(dx, dy)

	p1 := newVector(100, 64)
	p2 := newVector(20, 255)
	p3 := newVector(200, 200)
	cv.drawLine(p1, p2)
	cv.drawLine(p1, p3)
	cv.drawLine(p3, p2)
	//cv.drawPoint(p2)
	createPNG(cv.img)
}
