package utils

import (
	"math"
)

const (
	INFINITY = math.MaxFloat64
	PI       = math.Pi
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * PI / 180.0
}
