package main

type Vector struct {
	x int
	y int
}

func newVector(x int, y int) *Vector {
	return &Vector{
		x,
		y,
	}
}