package main

type Shape interface {
	GetPerimeter() float64
	GetArea() float64
}

func main() {
	shapes := []Shape{}

	shapes = append(shapes, Rectangle{...})
	shapes = append(shapes, Square{...})
	shapes = append(shapes, Triangle{...})

	for _, shape := range shapes {
		fmt.Printf("Perimeter: %d\n", shape.GetPerimeter())
		fmt.Printf("Area: %d\n", shape.GetArea())
	}
}
