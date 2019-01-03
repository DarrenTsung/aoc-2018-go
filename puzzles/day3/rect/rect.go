package rect

// Point is a point
type Point struct {
	X, Y int
}

// Rect is a rectangle
type Rect struct {
	X, Y, Width, Height int
}

// IterPoints calls the provided function for each point in the Rect
func (rect Rect) IterPoints(iter func(point Point)) {
	for x := rect.X; x < rect.X+rect.Width; x++ {
		for y := rect.Y; y < rect.Y+rect.Height; y++ {
			iter(Point{X: x, Y: y})
		}
	}
}

// ContainsPoint returns whether the x, y is contained in the Rect
func (rect Rect) ContainsPoint(x, y int) bool {
	return rect.X <= x &&
		x < rect.X+rect.Width &&
		rect.Y <= y &&
		y < rect.Y+rect.Height
}
