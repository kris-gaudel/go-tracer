package utils

import (
	"math"
	"math/rand"
)

const (
	INFINITY = math.MaxFloat64
	PI       = math.Pi
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * PI / 180.0
}

func RandomDouble() float64 {
	return rand.Float64()
}

func RandomDoubleRange(min, max float64) float64 {
	return min + (max-min)*RandomDouble()
}
