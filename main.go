package main

import (
	"flag"
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/vec3"
	"log"
	"time"
)

func setupWorld() hittable.HittableList {
	var world hittable.HittableList
	material_ground := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.8, Y: 0.8, Z: 0.0}}
	material_center := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.1, Y: 0.2, Z: 0.5}}
	material_left := hittable.Dielectric{Ir: 1.5}
	material_right := hittable.Metal{Albedo: vec3.Vec3{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 0.0}

	world.Append(hittable.Sphere{Center: vec3.Point3{X: 0, Y: -100.5, Z: -1}, Radius: 100, Mat: material_ground})
	world.Append(hittable.Sphere{Center: vec3.Point3{X: 0, Y: 0, Z: -1}, Radius: 0.5, Mat: material_center})
	world.Append(hittable.Sphere{Center: vec3.Point3{X: -1, Y: 0, Z: -1}, Radius: 0.5, Mat: material_left})
	world.Append(hittable.Sphere{Center: vec3.Point3{X: -1, Y: 0, Z: -1}, Radius: -0.4, Mat: material_left})
	world.Append(hittable.Sphere{Center: vec3.Point3{X: 1, Y: 0, Z: -1}, Radius: 0.5, Mat: material_right})

	return world
}

func setupCamera() camera.Camera {
	var cam camera.Camera
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 1200
	cam.SamplesPerPixel = 500
	cam.MaxDepth = 50

	cam.VFOV = 20.0
	cam.LookFrom = vec3.Point3{X: 13, Y: 2, Z: 3}
	cam.LookAt = vec3.Point3{X: 0, Y: 0, Z: 0}
	cam.ViewUp = vec3.Vec3{X: 0, Y: 1, Z: 0}
	cam.DefocusAngle = 0.6
	cam.FocusDistance = 10.0

	return cam
}

func main() {
	// Command line flags
	multiThread := flag.Bool("multi", true, "Use multi-threaded rendering")
	flag.Parse()

	// Setup scene
	world := setupWorld()
	cam := setupCamera()

	// Time the rendering
	start := time.Now()

	// Render based on flag
	if *multiThread {
		log.Printf("Starting multi-threaded render...")
		cam.RenderMulti(&world)
	} else {
		log.Printf("Starting single-threaded render...")
		cam.RenderSingle(&world)
	}

	// Calculate and display render time
	duration := time.Since(start)
	log.Printf("\nRendering completed in: %v", duration)
	log.Printf("Mode: %s", map[bool]string{true: "Multi-threaded", false: "Single-threaded"}[*multiThread])
}
