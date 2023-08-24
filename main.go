package main

import (
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/utils"
	"go-tracer/src/vec3"
	"math"
)

func main() {
	// World
	var world hittable.HittableList
	R := math.Cos(utils.PI / 4)
	material_left := hittable.Lambertian{Albedo: vec3.Vec3{X: 0, Y: 0, Z: 1}}
	material_right := hittable.Lambertian{Albedo: vec3.Vec3{X: 1, Y: 0, Z: 0}}
	sphereOne := hittable.Sphere{Center: vec3.Vec3{X: -R, Y: 0.0, Z: -1.0}, Radius: R, Mat: material_left}
	sphereTwo := hittable.Sphere{Center: vec3.Vec3{X: R, Y: 0.0, Z: -1.0}, Radius: R, Mat: material_right}
	world.Append(&sphereOne)
	world.Append(&sphereTwo)

	// Camera
	var cam camera.Camera
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50

	cam.VFOV = 90.0
	cam.Render(&world)

}
