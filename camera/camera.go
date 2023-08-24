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
	AspectRatio     float64
	ImageWidth      int
	ImageHeight     int
	SamplesPerPixel int
	MaxDepth        int
	VFOV            float64
	LookFrom        vec3.Point3
	LookAt          vec3.Point3
	ViewUp          vec3.Vec3
	Center          vec3.Point3
	Pixel00_loc     vec3.Point3
	PixelDeltaU     vec3.Vec3
	PixelDeltaV     vec3.Vec3
	U, V, W         vec3.Vec3
}

func (c *Camera) RayColor(r *vec3.Ray, depth int, world hittable.Hittable) vec3.Vec3 {
	var rec hittable.HitRecord

	if depth <= 0 {
		return vec3.Vec3{X: 0, Y: 0, Z: 0}
	}

	if world.Hit(r, interval.Interval{Min: 0.001, Max: utils.INFINITY}, &rec) {
		var scattered vec3.Ray
		var attenuation vec3.Vec3
		if (rec.Mat).Scatter(r, &rec, &attenuation, &scattered) {
			return *attenuation.MultiplyVec(c.RayColor(&scattered, depth-1, world))
		} else {
			return vec3.Vec3{X: 0, Y: 0, Z: 0}
		}
		// direction := rec.Normal.Add(*rec.Normal.RandomUnitVector())
		// computedValue := c.RayColor(&vec3.Ray{Origin: rec.P, Direction: direction}, depth-1, world).MultiplyFloat(0.1)
		// return *computedValue
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
			var pixel_color vec3.Vec3 = vec3.Vec3{X: 0, Y: 0, Z: 0}
			for sample := 0; sample < c.SamplesPerPixel; sample++ {
				r := c.GetRay(i, j)
				pixel_color.PlusEqual(c.RayColor(&r, c.MaxDepth, world))
			}
			fmt.Println(pixel_color.String((*c).SamplesPerPixel))
		}
	}
	log.Println("Done!")
}

func (c *Camera) GetRay(i, j int) vec3.Ray {
	pixel_center := c.Pixel00_loc.Add(*c.PixelDeltaU.MultiplyFloat(float64(i))).Add(*c.PixelDeltaV.MultiplyFloat(float64(j)))
	pixel_sample := pixel_center.Add(c.PixelSampleSquare())

	ray_origin := c.Center
	ray_direction := pixel_sample.Subtract(c.Center)

	return vec3.Ray{Origin: ray_origin, Direction: *ray_direction}
}

func (c *Camera) PixelSampleSquare() vec3.Vec3 {
	px := -0.5 + utils.RandomDouble()
	py := -0.5 + utils.RandomDouble()

	return c.PixelDeltaU.MultiplyFloat(px).Add(*c.PixelDeltaV.MultiplyFloat(py))
}

func (c *Camera) Initalize() {
	(*c).ImageWidth = 400
	(*c).ImageHeight = int(math.Max(float64((*c).ImageWidth)/(*c).AspectRatio, 1.0))

	(*c).Center = c.LookFrom

	focal_length := (c.LookFrom.Subtract(c.LookAt)).Length()
	theta := utils.DegreesToRadians(c.VFOV)
	h := math.Tan(theta / 2)
	viewport_height := 2.0 * h * focal_length
	viewport_width := viewport_height * (float64((*c).ImageWidth) / float64((*c).ImageHeight))

	(*c).W = *c.LookFrom.Subtract(c.LookAt).UnitVector()
	(*c).U = *c.ViewUp.Cross(c.W).UnitVector()
	(*c).V = *c.W.Cross(c.U)

	viewport_u := c.U.MultiplyFloat(viewport_width)
	viewport_v := c.V.Negate().MultiplyFloat(viewport_height)

	(*c).PixelDeltaU = *viewport_u.DivideFloat(float64(c.ImageWidth))
	(*c).PixelDeltaV = *viewport_v.DivideFloat(float64(c.ImageHeight))

	viewport_upper_left := c.Center.Subtract(*c.W.MultiplyFloat(focal_length)).Subtract(*viewport_u.DivideFloat(2)).Subtract(*viewport_v.DivideFloat(2))
	(*c).Pixel00_loc = viewport_upper_left.Add(*c.PixelDeltaU.DivideFloat(2)).Add(*c.PixelDeltaV.DivideFloat(2))
}
