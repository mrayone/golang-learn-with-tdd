package di

type Shape interface {
	Area() float64
	Width() float64
	Height() float64
}

type ShapeFactory interface {
	Make() Shape
}

type ShapeRetangleFactory struct{}

func (f ShapeRetangleFactory) Make() Shape {
	return NewRectangle(1, 2)
}

type Retangle struct {
	width, height float64
}

func NewRectangle(width, height float64) *Retangle {
	return &Retangle{width, height}
}

func (r *Retangle) Area() float64 {
	return r.height * r.width
}

func (r *Retangle) Width() float64 {
	return r.width
}

func (r *Retangle) Height() float64 {
	return r.height
}
