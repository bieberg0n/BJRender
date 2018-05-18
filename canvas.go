package main

import (
	"image"
	"image/color"
)

type Canvas struct {
	img *image.RGBA
	dx  int
	dy  int
}

func newCanvas(dx int, dy int) *Canvas {
	cv := &Canvas{
		image.NewRGBA(image.Rect(0, 0, dx, dy)),
		dx,
		dy,
	}
	for x := 0; x < dx; x++ {
		for y:= 0; y < dy; y ++ {
			cv.img.Set(x, y, color.RGBA{0,0,0,255})
		}
	}
	return cv
}

func (cv *Canvas) drawLine(p1 *Vector, p2 *Vector) {
	x1, y1, x2, y2 := p1.x, p1.y, p2.x, p2.y

	dx := x2 - x1
	dy := y2 - y1

	var vs []*Vector
	//logs(abs(dx), abs(dy))
	if abs(dx) >= abs(dy) {
		vs = lineVector(x1, x2, y1, y2, dx, dy, false)
	} else {
		vs = lineVector(y1, y2, x1, x2, dy, dx, true)
	}
	for _, v := range vs {
		logs(v.x, v.y)
		cv.drawPoint(v)
	}
}

func (cv *Canvas) drawPoint(p *Vector) {
	if 0 <= p.x &&
		p.x < cv.dx &&
		0 <= p.y &&
		p.y < cv.dy {
		cv.img.Set(p.x, p.y, color.RGBA{255, 255, 255, 255})
	}
}

func lineVector(x1, x2, y1, y2, dx, dy int, reversed bool) []*Vector {
	var (
		xmin int
		xmax int
		y float32
	)
	if x1 < x2 {
		xmin, xmax = x1, x2
		y = float32(y1)
	} else {
		xmin, xmax = x2, x1
		y = float32(y2)
	}

	var k float32
	if dx == 0 {
		k = 0
	} else {
		dyF := float32(dy)
		dxF := float32(dx)
		k = dyF / dxF
	}

	var vs []*Vector
	var v *Vector
	for x := xmin; x < xmax; x++ {
		y += k
		if reversed {
			v = newVector(int(y), x)
		} else {
			v = newVector(x, int(y))
		}
		vs = append(vs, v)
	}
	return vs
}