package geom

type rect struct {
	a int
	b int
}

type circ struct {
	a int
}

func NewRect(a, b int) rect {
	return rect{a, b}
}
func NewCirc(a int) circ {
	return circ{a}
}

func (r rect) FindS() int {
	return r.a * r.b
}

func (c circ) FindS() int {
	return c.a * c.a * 3
}

//+++++++++++++++++++

func FindA(a,b int) int {
	return a+b
}