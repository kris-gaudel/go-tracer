package camera

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

type Camera struct {
	AspectRatio float64
	ImageWidth  int
	ImageHeight int
	Center      vec3.Point3
	Pixel00_loc vec3.Point3
	PixelDeltaU vec3.Vec3
	PixelDeltaV vec3.Vec3
}

func (c *Camera) RayColor(r *vec3.Ray, world hittable.Hittable) vec3.Vec3 {
	var rec hittable.HitRecord
	if world.Hit(r, interval.Interval{Min: 0, Max: utils.INFINITY}, &rec) {
		computedValue := rec.Normal.Add(vec3.Vec3{X: 1, Y: 1, Z: 1}).MultiplyFloat(0.5)
		return *computedValue
	}

	unit_direction := r.GetDirection().UnitVector()
	a := 0.5 * (unit_direction.GetY() + 1.0)
	startValue := vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	endValue := vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	computedValue := startValue.MultiplyFloat(1.0 - a).Add(*endValue.MultiplyFloat(a))
	return computedValue
}

func (c *Camera) Render(world hittable.Hittable) {
	c.Initalize()
	fmt.Println("P3")
	fmt.Println(strconv.Itoa(c.ImageWidth) + " " + strconv.Itoa(c.ImageHeight))
	fmt.Println("255")
	for j := 0; j < c.ImageHeight; j++ {
		log.Println("Scanlines remaining: " + strconv.Itoa(c.ImageHeight-j))
		for i := 0; i < c.ImageWidth; i++ {
			pixel_center := c.Pixel00_loc.Add(*c.PixelDeltaU.MultiplyFloat(float64(i))).Add(*c.PixelDeltaV.MultiplyFloat(float64(j)))
			ray_direction := pixel_center.Subtract(c.Center)
			r := vec3.Ray{Origin: c.Center, Direction: *ray_direction}
			pixel_color := c.RayColor(&r, world)

			fmt.Println(pixel_color.String())
		}
	}
	log.Println("Done!")
}

func (c *Camera) Initalize() {
	// Image
	(*c).AspectRatio = 16.0 / 9.0
	(*c).ImageWidth = 400
	(*c).ImageHeight = int(math.Max(float64((*c).ImageWidth)/(*c).AspectRatio, 1.0))

	// Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64((*c).ImageWidth) / float64((*c).ImageHeight))
	camera_center := vec3.Point3{X: 0, Y: 0, Z: 0}

	viewport_u := vec3.Vec3{X: viewport_width, Y: 0, Z: 0}
	viewport_v := vec3.Vec3{X: 0, Y: -viewport_height, Z: 0}

	(*c).PixelDeltaU = *viewport_u.DivideFloat(float64((*c).ImageWidth))
	(*c).PixelDeltaV = *viewport_v.DivideFloat(float64((*c).ImageHeight))

	viewport_upper_left := camera_center.Subtract(vec3.Vec3{X: 0, Y: 0, Z: focal_length}).Subtract(*viewport_u.DivideFloat(2)).Subtract(*viewport_v.DivideFloat(2))
	(*c).Pixel00_loc = viewport_upper_left.Add(*c.PixelDeltaU.DivideFloat(2)).Add(*c.PixelDeltaV.DivideFloat(2))
}
