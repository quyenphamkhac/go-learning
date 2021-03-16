package main

// Point ...
type Point struct {
	X int
	Y int
}

// Scale ...
func (p Point) Scale(n int) {
	p.X = p.X * n
	p.Y = p.Y * n
}

// Scale1 ...
func (p *Point) Scale1(n int) {
	p.X = p.X * n
	p.Y = p.Y * n
}

// ValueReceiver ...
func ValueReceiver(p Point, ratio int) Point {
	p.X = p.X * ratio
	p.Y = p.Y * ratio
	return p
}

// PointerReceiver ...
func PointerReceiver(p *Point, ratio int) {
	p.X = p.X * ratio
	p.Y = p.Y * ratio
}

func main() {
}
