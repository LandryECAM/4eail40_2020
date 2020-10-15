package main

import "fmt"

// Implement types Rectangle, Circle and Shape

type Shape interface {
	// String() string
	fmt.Stringer
}

type Rectangle struct {
	W int
	L int
}

type Circle struct {
	R int
}

func (rectangle Rectangle) String() string {
	return fmt.Sprintf("Square of width %d and length %d", rectangle.W, rectangle.L)
}

func (circle Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", circle.R)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
