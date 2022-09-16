package raytrace

type Ray struct {
	orig Point3
	dir  Vec3
}

func (r Ray) Origin() Point3 {
	return r.orig
}

func (r Ray) Direction() Vec3 {
	return r.dir
}

func (r Ray) At(t float64) Point3 {
	u := Vec3(r.orig)
	v := r.dir.MulBy(t)
	return Point3(u.Add(v))
}
