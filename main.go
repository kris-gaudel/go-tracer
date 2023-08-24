package main

import (
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/utils"
	"go-tracer/src/vec3"
)

func main() {
	// World
	var world hittable.HittableList
	ground_material := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}
	sphereOne := hittable.Sphere{Center: vec3.Vec3{X: 0, Y: -1000, Z: 0}, Radius: 1000, Mat: ground_material}
	world.Append(&sphereOne)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := utils.RandomDouble()
			center := vec3.Point3{X: float64(a) + 0.9*utils.RandomDouble(), Y: 0.2, Z: float64(b) + 0.9*utils.RandomDouble()}

			if ((center.Subtract(vec3.Point3{X: 4, Y: 0.2, Z: 0})).Length()) > 0.9 {
				var sphere_material hittable.Material
				if choose_mat < 0.8 {
					albedo := vec3.Vec3{X: utils.RandomDouble(), Y: utils.RandomDouble(), Z: utils.RandomDouble()}
					sphere_material = hittable.Lambertian{Albedo: albedo}
					world.Append(&hittable.Sphere{Center: center, Radius: 0.2, Mat: sphere_material})
				} else if choose_mat < 0.95 {
					albedo := vec3.Vec3{X: utils.RandomDoubleRange(0.5, 1), Y: utils.RandomDoubleRange(0.5, 1), Z: utils.RandomDoubleRange(0.5, 1)}
					fuzz := utils.RandomDoubleRange(0, 0.5)
					sphere_material = hittable.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Append(&hittable.Sphere{Center: center, Radius: 0.2, Mat: sphere_material})
				} else {
					sphere_material = hittable.Dielectric{Ir: 1.5}
					world.Append(&hittable.Sphere{Center: center, Radius: 0.2, Mat: sphere_material})
				}
			}
		}
	}

	materialOne := hittable.Dielectric{Ir: 1.5}
	world.Append(&hittable.Sphere{Center: vec3.Point3{X: 0, Y: 1, Z: 0}, Radius: 1.0, Mat: materialOne})

	materialTwo := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.4, Y: 0.2, Z: 0.1}}
	world.Append(&hittable.Sphere{Center: vec3.Point3{X: -4, Y: 1, Z: 0}, Radius: 1.0, Mat: materialTwo})

	materialThree := hittable.Metal{Albedo: vec3.Vec3{X: 0.7, Y: 0.6, Z: 0.5}, Fuzz: 0.0}
	world.Append(&hittable.Sphere{Center: vec3.Point3{X: 4, Y: 1, Z: 0}, Radius: 1.0, Mat: materialThree})

	// Camera
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

	// Default camera settings
	// cam.AspectRatio = 16.0 / 9.0
	// cam.ImageWidth = 400
	// cam.SamplesPerPixel = 100
	// cam.MaxDepth = 50
	// cam.LookFrom = vec3.Point3{X: 0, Y: 0, Z: -1}
	// cam.LookAt = vec3.Point3{X: 0, Y: 0, Z: 0}
	// cam.ViewUp = vec3.Vec3{X: 0, Y: 1, Z: 0}
	// cam.VFOV = 90.0
	// cam.DefocusAngle = 0.0
	// cam.FocusDistance = 10.0
	cam.Render(&world)

}
