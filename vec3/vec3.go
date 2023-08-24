package vec3

import (
	"fmt"
	"go-tracer/src/interval"
	"go-tracer/src/utils"
	"math"
)

const (
	COLOR_MAX     = 255.999
	COLOR_MAX_INT = 256
)

type Vec3 struct {
	X, Y, Z float64
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

func (v Vec3) IndexAt(i int) float64 {
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
	(*v).X += v2.X
	(*v).Y += v2.Y
	(*v).Z += v2.Z
	return v
}

func (v *Vec3) TimesEqual(t float64) *Vec3 {
	(*v).X *= t
	(*v).Y *= t
	(*v).Z *= t
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

func (v Vec3) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func (v Vec3) LinearToGamma(linear_component float64) float64 {
	return math.Sqrt(linear_component)
}

// Simulating the << overload, but writing our own String() method
func (v Vec3) String(samples_per_pixel int) string {
	r := v.GetX()
	g := v.GetY()
	b := v.GetZ()

	scale := 1.0 / float64(samples_per_pixel)
	r *= scale
	g *= scale
	b *= scale

	r = v.LinearToGamma(r)
	g = v.LinearToGamma(g)
	b = v.LinearToGamma(b)

	var intensity interval.Interval = interval.Interval{Min: 0.000, Max: 0.999}
	return fmt.Sprintf("%d %d %d", int(COLOR_MAX_INT*intensity.Clamp(r)),
		int(COLOR_MAX_INT*intensity.Clamp(g)),
		int(COLOR_MAX_INT*intensity.Clamp(b)))
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

func (v Vec3) RandomInUnitDisk() Vec3 {
	p := Vec3{X: utils.RandomDoubleRange(-1, 1), Y: utils.RandomDoubleRange(-1, 1), Z: 0}
	if p.LengthSquared() < 1.0 {
		return p
	}
	return p.RandomInUnitDisk()
}

func (v Vec3) Random() *Vec3 {
	return &Vec3{X: utils.RandomDouble(), Y: utils.RandomDouble(), Z: utils.RandomDouble()}
}

func (v Vec3) RandomRange(min, max float64) *Vec3 {
	return &Vec3{X: utils.RandomDoubleRange(min, max), Y: utils.RandomDoubleRange(min, max), Z: utils.RandomDoubleRange(min, max)}
}

func (v Vec3) RandomInUnitSphere() *Vec3 {
	p := v.RandomRange(-1, 1)
	if p.LengthSquared() < 1 {
		return v.RandomInUnitSphere()
	}
	return p
}

func (v Vec3) RandomUnitVector() *Vec3 {
	return v.RandomInUnitSphere().UnitVector()
}

func (v Vec3) RandomOnHemiSphere(normal *Vec3) *Vec3 {
	on_unit_sphere := v.RandomUnitVector()
	if on_unit_sphere.Dot(*normal) > 0.0 {
		return on_unit_sphere
	} else {
		return on_unit_sphere.MultiplyFloat(-1.0)
	}
}

func (v Vec3) Reflect(n *Vec3) Vec3 {
	return *v.Subtract(*n.MultiplyFloat(2 * v.Dot(*n)))
}

func (v Vec3) Refract(uv *Vec3, n *Vec3, etai_over_etat float64) Vec3 {
	cos_theta := math.Min(uv.Negate().Dot(*n), 1.0)
	r_out_perp := uv.Add(*n.MultiplyFloat(cos_theta)).MultiplyFloat(etai_over_etat)
	r_out_parallel := (*n).MultiplyFloat(-math.Sqrt(math.Abs(1.0 - r_out_perp.LengthSquared())))
	return r_out_perp.Add(*r_out_parallel)
}
