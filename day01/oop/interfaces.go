package oop

// ShaperArea is the interface for shapes
type ShaperArea interface {
	Area() int
}

// ShaperPerimeter is the interface for shapes
type ShaperPerimeter interface {
	Perimeter() int
}

// Rectangle is an implementation
type Rectangle struct {
	Length int
	Width  int
}

// Area is Rectangle area implementation
func (r Rectangle) Area() int {
	return r.Length * r.Width
}

// Perimeter is Rectangle perimeter implementation
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

// Square is an implementation
type Square struct {
	Side int
}

// Area is Square area implementation
func (sq Square) Area() int {
	return sq.Side * sq.Side
}

// Perimeter is Square perimeter implementation
func (sq Square) Perimeter() int {
	return 4 * sq.Side
}
