package raytrace

type Ray struct {
	Orig Point3
	Dir  Vec3
}

func (r Ray) Origin() Point3 {
	return r.Orig
}

func (r Ray) Direction() Vec3 {
	return r.Dir
}

func (r Ray) At(t float64) Point3 {
	u := r.Orig.Vec3
	v := r.Dir.MulBy(t)

	return Point3{u.Add(v)}
}
