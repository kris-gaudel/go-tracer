package hittable

import (
	"go-tracer/src/interval"
	"go-tracer/src/vec3"
	"math"
)

type Material interface {
	Scatter(r_in *vec3.Ray, rec *HitRecord, attenuation *vec3.Vec3, scattered *vec3.Ray) bool
}

type Lambertian struct {
	Albedo vec3.Vec3
}

func (l Lambertian) Scatter(r_in *vec3.Ray, rec *HitRecord, attenuation *vec3.Vec3, scattered *vec3.Ray) bool {
	scatter_direction := rec.Normal.Add(*rec.Normal.RandomUnitVector())
	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}

	(*scattered) = vec3.Ray{Origin: rec.P, Direction: scatter_direction}
	(*attenuation) = l.Albedo
	return true
}

type Metal struct {
	Albedo vec3.Vec3
}

func (m Metal) Scatter(r_in *vec3.Ray, rec *HitRecord, attenuation *vec3.Vec3, scattered *vec3.Ray) bool {
	reflected := r_in.GetDirection().UnitVector().Reflect(&rec.Normal)
	(*scattered) = vec3.Ray{Origin: rec.P, Direction: reflected}
	(*attenuation) = m.Albedo
	return (*scattered).GetDirection().Dot(rec.Normal) > 0
}

type HitRecord struct {
	P         vec3.Point3
	Normal    vec3.Vec3
	Mat       Material
	T         float64
	FrontFace bool
}

type Hittable interface {
	Hit(r *vec3.Ray, ray_t interval.Interval, rec *HitRecord) bool
}

type HittableList struct {
	Objects []Hittable
}

func (hl *HittableList) Clear() {
	hl.Objects = nil
}

func (hl *HittableList) Append(object Hittable) {
	if hl.Objects == nil {
		hl.Objects = make([]Hittable, 0)
	}
	hl.Objects = append(hl.Objects, object)
}

func (hl HittableList) Hit(r *vec3.Ray, ray_t interval.Interval, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := ray_t.Max

	for _, object := range hl.Objects {
		if object.Hit(r, interval.Interval{Min: ray_t.Min, Max: closestSoFar}, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}
	return hitAnything
}

// Define our shapes here
type Sphere struct {
	Hittable
	Center vec3.Point3
	Radius float64
	Mat    Material
}

func (hr *HitRecord) SetFaceNormal(ray *vec3.Ray, outwardNormal *vec3.Vec3) {
	hr.FrontFace = ray.GetDirection().Dot(*outwardNormal) < 0
	if hr.FrontFace {
		(*hr).Normal = *outwardNormal
	} else {
		(*hr).Normal = *(*outwardNormal).MultiplyFloat(-1)
	}
}

func (s Sphere) Hit(r *vec3.Ray, ray_t interval.Interval, rec *HitRecord) bool {
	oc := r.GetOrigin().Subtract(s.Center)
	a := r.GetDirection().LengthSquared()
	half_b := oc.Dot(r.GetDirection())
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (-half_b - sqrtd) / a
	if !ray_t.Surrounds(root) {
		root = (-half_b + sqrtd) / a
		if !ray_t.Surrounds(root) {
			return false
		}
	}

	(*rec).T = root
	(*rec).P = r.At(rec.T)
	outwardNormal := *(*rec.P.Subtract(s.Center)).DivideFloat(s.Radius)
	(*rec).SetFaceNormal(r, &outwardNormal)
	(*rec).Mat = s.Mat

	return true
}
