package hittable

import (
	"go-tracer/src/vec3"
	"testing"
)

func TestLambertianScatter(t *testing.T) {
	lambertian := Lambertian{Albedo: vec3.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}
	r_in := &vec3.Ray{}
	rec := &HitRecord{}
	attenuation := &vec3.Vec3{}
	scattered := &vec3.Ray{}

	result := lambertian.Scatter(r_in, rec, attenuation, scattered)

	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestMetalScatter(t *testing.T) {
	metal := Metal{Albedo: vec3.Vec3{X: 0.7, Y: 0.7, Z: 0.7}, Fuzz: 0.1}
	r_in := &vec3.Ray{}
	rec := &HitRecord{}
	attenuation := &vec3.Vec3{}
	scattered := &vec3.Ray{}

	result := metal.Scatter(r_in, rec, attenuation, scattered)

	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestDielectricScatter(t *testing.T) {
	dielectric := Dielectric{Ir: 1.5}
	r_in := &vec3.Ray{}
	rec := &HitRecord{}
	attenuation := &vec3.Vec3{}
	scattered := &vec3.Ray{}

	result := dielectric.Scatter(r_in, rec, attenuation, scattered)

	if !result {
		t.Errorf("Expected true, but got false")
	}
}
