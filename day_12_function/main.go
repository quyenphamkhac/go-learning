package main

// Vertex struct
type Vertex struct {
	X int
	Y int
}

// Pic func
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}

func multier() func(int) int {
	mul := 1
	return func(x int) int {
		mul = mul * x
		return mul
	}
}

func fibonacci() func() int {
	f0 := 0
	f1 := 0
	return func() int {
		f1, f0 = f0, f1
		f1 = f0 + f1
		return f1
	}
}

// Scale method
func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc func
func ScaleFunc(v *Vertex, f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
}
