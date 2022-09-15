package vec

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

// Type alias for vec3
type Point3 Vec3 // 3D Point
type Color Vec3  // RBG Color

func MakeVec3(e0 float64, e1 float64, e2 float64) Vec3 {
	return Vec3{e: [3]float64{e0, e1, e2}}
}
func (v Vec3) X() float64 {
	return v.e[0]
}

func (v Vec3) Y() float64 {
	return v.e[1]
}

func (v Vec3) Z() float64 {
	return v.e[2]
}

func (v Vec3) Get(i int) float64 {
	return v.e[i]
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v Vec3) LenSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.e[0], v.e[1], v.e[2])
}

func (v Vec3) Add(u Vec3) Vec3 {
	return MakeVec3(u.e[0]+v.e[0], u.e[1]+v.e[1], u.e[2]+v.e[2])
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return MakeVec3(u.e[0]-v.e[0], u.e[1]-v.e[1], u.e[2]-v.e[2])
}

func (v Vec3) Mul(u Vec3) Vec3 {
	return MakeVec3(u.e[0]*v.e[0], u.e[1]*v.e[1], u.e[2]*v.e[2])
}

func (v Vec3) MulBy(t float64) Vec3 {
	return MakeVec3(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func (v Vec3) DivBy(t float64) Vec3 {
	return v.MulBy(1 / t)
}

func (v Vec3) Dot(u Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return MakeVec3(u.e[1]*v.e[2]-u.e[2]*v.e[1], u.e[2]*v.e[0]-u.e[0]*v.e[2], u.e[0]*v.e[1]-u.e[1]*v.e[0])
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivBy(v.Len())
}
