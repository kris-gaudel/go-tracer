package main

import (
	"fmt"
	"go-tracer/src/vec3"
	"log"
	"math"
	"strconv"
)

func HitSphere(center *vec3.Point3, radius float64, r *vec3.Ray) bool {
	oc := r.GetOrigin().Subtract(*center)
	a := r.GetDirection().Dot(r.GetDirection())
	b := 2.0 * oc.Dot(r.GetDirection())
	c := oc.Dot(*oc) - radius*radius
	discriminant := b*b - 4*a*c

	return (discriminant >= 0)
}

func RayColor(r *vec3.Ray) vec3.Color {
	if HitSphere(&vec3.Point3{X: 0, Y: 0, Z: -1}, 0.5, r) {
		computedValue := vec3.Vec3{X: 1.0, Y: 0.0, Z: 0.0}
		return *computedValue.ConvertToRGB()
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
			ray := vec3.Ray{Origin: camera_center, Direction: *ray_direction}
			pixel_color := RayColor(&ray)

			fmt.Println(pixel_color.String())
		}
	}
	log.Println("Done!")
}
