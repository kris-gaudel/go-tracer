package main

import (
	"fmt"
	"go-tracer/src/hittable"
	"go-tracer/src/interval"
	"go-tracer/src/utils"
	"go-tracer/src/vec3"
	"log"
	"math"
	"strconv"
)

func RayColor(r *vec3.Ray, world hittable.Hittable) vec3.Color {
	var rec hittable.HitRecord
	if world.Hit(r, interval.Interval{Min: 0, Max: utils.INFINITY}, &rec) {
		computed_value := rec.Normal.Add(vec3.Vec3{X: 1, Y: 1, Z: 1}).MultiplyFloat(0.5)
		return *computed_value.ConvertToRGB()
	}

	unit_direction := r.GetDirection().UnitVector()
	a := 0.5 * (unit_direction.GetY() + 1.0)
	startValue := vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	endValue := vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	computedValue := startValue.MultiplyFloat(1.0 - a).Add(*endValue.MultiplyFloat(a))
	return *computedValue.ConvertToRGB()
}

func main() {
	// Image
	aspect_ratio := 16.0 / 9.0
	var image_width int = 400
	var image_height int = int(math.Max(float64(image_width)/aspect_ratio, 1.0))

	// World
	var world hittable.HittableList
	sphereOne := hittable.Sphere{Center: vec3.Point3{X: 0, Y: 0, Z: -1}, Radius: 0.5}
	sphereTwo := hittable.Sphere{Center: vec3.Point3{X: 0, Y: -100.5, Z: -1}, Radius: 100}
	world.Append(sphereOne)
	world.Append(sphereTwo)

	// Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64(image_width) / float64(image_height))
	camera_center := vec3.Point3{X: 0, Y: 0, Z: 0}

	viewport_u := vec3.Vec3{X: viewport_width, Y: 0, Z: 0}
	viewport_v := vec3.Vec3{X: 0, Y: -viewport_height, Z: 0}

	pixel_delta_u := viewport_u.DivideFloat(float64(image_width))
	pixel_delta_v := viewport_v.DivideFloat(float64(image_height))

	viewport_upper_left := camera_center.Subtract(vec3.Vec3{X: 0, Y: 0, Z: focal_length}).Subtract(*viewport_u.DivideFloat(2)).Subtract(*viewport_v.DivideFloat(2))
	pixel00_loc := viewport_upper_left.Add(*pixel_delta_u.DivideFloat(2)).Add(*pixel_delta_v.DivideFloat(2))

	// Render
	fmt.Println("P3")
	fmt.Println(strconv.Itoa(image_width) + " " + strconv.Itoa(image_height))
	fmt.Println("255")
	for j := 0; j < image_height; j++ {
		log.Println("Scanlines remaining: " + strconv.Itoa(image_height-j))
		for i := 0; i < image_width; i++ {
			pixel_center := pixel00_loc.Add(*pixel_delta_u.MultiplyFloat(float64(i))).Add(*pixel_delta_v.MultiplyFloat(float64(j)))
			ray_direction := pixel_center.Subtract(camera_center)
			r := vec3.Ray{Origin: camera_center, Direction: *ray_direction}
			pixel_color := RayColor(&r, world)

			fmt.Println(pixel_color.String())
		}
	}
	log.Println("Done!")
}
