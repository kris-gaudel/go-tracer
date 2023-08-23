package main

import (
	"go-tracer/src/camera"
	"go-tracer/src/hittable"
	"go-tracer/src/interval"
	"go-tracer/src/utils"
	"go-tracer/src/vec3"
)

func RayColor(r *vec3.Ray, world hittable.Hittable) vec3.Vec3 {
	var rec hittable.HitRecord
	if world.Hit(r, interval.Interval{Min: 0, Max: utils.INFINITY}, &rec) {
		computed_value := rec.Normal.Add(vec3.Vec3{X: 1, Y: 1, Z: 1}).MultiplyFloat(0.5)
		return *computed_value
	}

	unit_direction := r.GetDirection().UnitVector()
	a := 0.5 * (unit_direction.GetY() + 1.0)
	startValue := vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	endValue := vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	computedValue := startValue.MultiplyFloat(1.0 - a).Add(*endValue.MultiplyFloat(a))
	return computedValue
}

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
	cam.Render(&world)

}
