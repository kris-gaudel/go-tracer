package main

import (
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/vec3"
)

func main() {
	// World
	var world hittable.HittableList
	sphereOne := hittable.Sphere{Center: vec3.Point3{X: 0, Y: 0, Z: -1}, Radius: 0.5}
	sphereTwo := hittable.Sphere{Center: vec3.Point3{X: 0, Y: -100.5, Z: -1}, Radius: 100}
	world.Append(sphereOne)
	world.Append(sphereTwo)

	// Camera
	var cam camera.Camera
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50
	cam.Render(&world)

}
