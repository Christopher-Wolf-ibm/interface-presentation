package main

type Polygon interface {
	GetPerimeter() float64
	GetArea() float64
}

func main() {
	polygons := []Polygon{}

	polygons = append(polygons, Rectangle{...})
	polygons = append(polygons, Square{...})
	polygons = append(polygons, Triangle{...})

	for _, polygon := range polygons {
		fmt.Printf("Perimeter: %d\n", polygon.GetPerimeter())
		fmt.Printf("Area: %d\n\n", polygon.GetArea())
	}
}
