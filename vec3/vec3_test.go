package vec3

import (
	"math"
	"testing"
)

const EPSILON = 1e-7

func almostEqual(t *testing.T, got, want Vec3, msg string) {
	if math.Abs(got.X-want.X) > EPSILON ||
		math.Abs(got.Y-want.Y) > EPSILON ||
		math.Abs(got.Z-want.Z) > EPSILON {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}

func floatAlmostEqual(t *testing.T, got, want float64, msg string) {
	if math.Abs(got-want) > EPSILON {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}

func TestVec3BasicOperations(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*testing.T)
	}{
		{"Addition", func(t *testing.T) {
			v1 := Vec3{1, 2, 3}
			v2 := Vec3{4, 5, 6}
			result := v1.Add(v2)
			want := Vec3{5, 7, 9}
			almostEqual(t, result, want, "Vector addition")
		}},

		{"Subtraction", func(t *testing.T) {
			v1 := Vec3{4, 5, 6}
			v2 := Vec3{1, 2, 3}
			result := *v1.Subtract(v2)
			want := Vec3{3, 3, 3}
			almostEqual(t, result, want, "Vector subtraction")
		}},

		{"Scalar Multiplication", func(t *testing.T) {
			v := Vec3{1, 2, 3}
			result := *v.MultiplyFloat(2.0)
			want := Vec3{2, 4, 6}
			almostEqual(t, result, want, "Scalar multiplication")
		}},

		{"Length", func(t *testing.T) {
			v := Vec3{3, 4, 0}
			got := v.Length()
			want := 5.0
			floatAlmostEqual(t, got, want, "Vector length")
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.fn)
	}
}

func TestVec3DotProduct(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{4, 5, 6}
	got := v1.Dot(v2)
	want := 32.0 // (1*4 + 2*5 + 3*6)
	floatAlmostEqual(t, got, want, "Dot product")
}

func TestVec3CrossProduct(t *testing.T) {
	v1 := Vec3{1, 0, 0}
	v2 := Vec3{0, 1, 0}
	result := *v1.Cross(v2)
	want := Vec3{0, 0, 1}
	almostEqual(t, result, want, "Cross product")
}

func TestVec3UnitVector(t *testing.T) {
	v := Vec3{3, 4, 0}
	result := *v.UnitVector()
	want := Vec3{0.6, 0.8, 0}
	almostEqual(t, result, want, "Unit vector")
}

func TestRayOperations(t *testing.T) {
	origin := Point3{1, 2, 3}
	direction := Vec3{0, 1, 0}
	ray := Ray{Origin: origin, Direction: direction}

	t.Run("Ray At", func(t *testing.T) {
		point := ray.At(2.0)
		want := Vec3{1, 4, 3}
		almostEqual(t, point, want, "Ray position at t=2")
	})
}

func TestVec3NearZero(t *testing.T) {
	tests := []struct {
		name string
		vec  Vec3
		want bool
	}{
		{
			name: "Near zero vector",
			vec:  Vec3{1e-9, 1e-9, 1e-9},
			want: true,
		},
		{
			name: "Non-zero vector",
			vec:  Vec3{0.1, 0.1, 0.1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.vec.NearZero()
			if got != tt.want {
				t.Errorf("NearZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3Reflect(t *testing.T) {
	v := Vec3{1, -1, 0}
	n := Vec3{0, 1, 0}
	result := v.Reflect(&n)
	want := Vec3{1, 1, 0}
	almostEqual(t, result, want, "Vector reflection")
}
