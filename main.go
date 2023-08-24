package main

import (
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/vec3"
)

func main() {
	// World
	var world hittable.HittableList
	material_ground := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.8, Y: 0.8, Z: 0.0}}
	material_center := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.1, Y: 0.2, Z: 0.5}}
	material_left := hittable.Dielectric{Ir: 1.5}
	material_right := hittable.Metal{Albedo: vec3.Vec3{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 0.0}

	sphereOne := hittable.Sphere{Center: vec3.Point3{X: 0.0, Y: -100.5, Z: -1.0}, Radius: 100.0, Mat: material_ground}
	sphereTwo := hittable.Sphere{Center: vec3.Point3{X: 0.0, Y: 0.0, Z: -1.0}, Radius: 0.5, Mat: material_center}
	sphereThree := hittable.Sphere{Center: vec3.Point3{X: -1.0, Y: 0.0, Z: -1.0}, Radius: 0.5, Mat: material_left}
	sphereFour := hittable.Sphere{Center: vec3.Point3{X: -1.0, Y: 0.0, Z: -1.0}, Radius: -0.4, Mat: material_left}
	sphereFive := hittable.Sphere{Center: vec3.Point3{X: 1.0, Y: 0.0, Z: -1.0}, Radius: 0.5, Mat: material_right}

	world.Append(&sphereOne)
	world.Append(&sphereTwo)
	world.Append(&sphereThree)
	world.Append(&sphereFour)
	world.Append(&sphereFive)

	// Camera
	var cam camera.Camera
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50

	cam.VFOV = 20.0
	cam.LookFrom = vec3.Point3{X: -2, Y: 2, Z: 1}
	cam.LookAt = vec3.Point3{X: 0, Y: 0, Z: -1}
	cam.ViewUp = vec3.Vec3{X: 0, Y: 1, Z: 0}
	cam.DefocusAngle = 10.0
	cam.FocusDistance = 3.4

	// Default camera settings
	// cam.LookFrom = vec3.Point3{X: 0, Y: 0, Z: -1}
	// cam.LookAt = vec3.Point3{X: 0, Y: 0, Z: 0}
	// cam.ViewUp = vec3.Vec3{X: 0, Y: 1, Z: 0}
	// cam.VFOV = 90.0
	// cam.DefocusAngle = 0.0
	// cam.FocusDistance = 10.0
	cam.Render(&world)

}
