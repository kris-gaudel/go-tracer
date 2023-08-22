package vec3

import (
	"fmt"
	"math"
)

const (
	COLOR_MAX = 255.999
)

type Vec3 struct {
	X, Y, Z float64
}

type Color struct {
	ColorX, ColorY, ColorZ int
}

type Point3 = Vec3

type Ray struct {
	Origin    Point3
	Direction Vec3
}

// Ray functions

func (r Ray) GetOrigin() Point3 {
	return r.Origin
}

func (r Ray) GetDirection() Vec3 {
	return r.Direction
}

func (r Ray) At(t float64) Vec3 {

	return r.GetOrigin().Add(*r.GetDirection().MultiplyFloat(t))
}

// Color functions
func (v Vec3) ConvertToRGB() *Color {
	r := int(COLOR_MAX * v.GetX())
	g := int(COLOR_MAX * v.GetY())
	b := int(COLOR_MAX * v.GetZ())
	return &Color{r, g, b}
}

func (c Color) String() string {
	return fmt.Sprintf("%d %d %d", c.ColorX, c.ColorY, c.ColorZ)
}

// Defining "class" methods for Vec3
func (v Vec3) GetX() float64 {
	return v.X
}

func (v Vec3) GetY() float64 {
	return v.Y
}

func (v Vec3) GetZ() float64 {
	return v.Z
}

// Go doesn't have operator overloading, so we have to define these methods
func (v Vec3) Negate() Vec3 {
	// Simulate the - operator overload (negation)
	return Vec3{-(v.X), -(v.Y), -(v.Z)}
}

func (v Vec3) IndeXAt(i int) float64 {
	// Simulate the [] operator overload
	if i == 0 {
		return v.X
	} else if i == 1 {
		return v.Y
	} else if i == 2 {
		return v.Z
	} else {
		return 0
	}
}

func (v *Vec3) PlusEqual(v2 Vec3) *Vec3 {
	v.Z += v2.Z
	v.Y += v2.Y
	v.Z += v2.Z
	return v
}

func (v *Vec3) TimesEqual(t float64) *Vec3 {
	v.X *= t
	v.Y *= t
	v.Z *= t
	return v
}

func (v Vec3) DivideEqual(t float64) *Vec3 {
	return v.TimesEqual(1 / t)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Simulating the << overload, but writing our own String() method
func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) Subtract(v2 Vec3) *Vec3 {
	return &Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3) MultiplyVec(v2 Vec3) *Vec3 {
	return &Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vec3) MultiplyFloat(t float64) *Vec3 {
	return &Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) DivideFloat(t float64) *Vec3 {
	return v.MultiplyFloat(1 / t)
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) *Vec3 {
	return &Vec3{v.Y*v2.Z - v.Z*v2.Y, v.Z*v2.X - v.X*v2.Z, v.X*v2.Y - v.Y*v2.X}
}

func (v Vec3) UnitVector() *Vec3 {
	return v.DivideFloat(v.Length())
}
