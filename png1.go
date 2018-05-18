package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

const (
	dx = 512
	dy = 512
	N  = 64
	PI = math.Pi
)

func circleSDF(x float32, y float32, cx float32, cy float32, r float32) float32 {
	ux := x - cx
	uy := y - cy
	return float32(math.Sqrt(float64(ux*ux+uy*uy))) - r
}

func trace(ox float32, oy float32, dx float32, dy float32) float32 {
	maxStep := 10
	var maxDistance float32 = 2.0
	var epsilon float32 = 1e-6

	var t float32 // t 为步进的距离
	var sd float32
	for i := 0; i < maxStep && t < maxDistance; i++ {
		// 光源中心为 (sourceX, sourceY)
		// 沿单位向量 (dx, dy) 方向前进，t 表示前进的距离
		sd = circleSDF(ox+dx*t, oy+dy*t, 0.5, 0.5, 0.1)

		// 此时已到达发光的圆形表面
		if sd < epsilon {
			return 2.0
		}
		// 继续增加步进的距离
		t += sd
	}
	return 0.0
}

func simple(x float32, y float32) float32 {
	var (
		theta float32
		sum   float32
	)
	for i := 0; i < N; i++ {
		// theta = PI * 2 * rand.Float32() // 随机采样
		// theta = PI * 2 * float32(i) / N // 分层采样（stratified sampling）
		theta = PI * 2 * (float32(i) + rand.Float32()) / N
		theta64 := float64(theta)
		sum += trace(x, y, float32(math.Cos(theta64)), float32(math.Sin(theta64)))
	}
	return sum / N
}

func createIMG1() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			c := min([]float32{simple(float32(x)/dx, float32(y)/dy) * 255, 255.0})
			uint8c := uint8(c)
			img.Set(x, y, color.RGBA{uint8c, uint8c, uint8c, 255})
		}
	}
	return img
}

