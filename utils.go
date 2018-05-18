package main

import (
	"fmt"
)

func logs(args ...interface{}) {
	fmt.Println(args...)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func min(arr []float32) float32 {
	n := arr[0]
	for i := 1; i < len(arr); i++ {
		if n > arr[i] {
			n = arr[i]
		}
	}
	return n
}

func pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j % 255)
		}
	}
	return pic
}
