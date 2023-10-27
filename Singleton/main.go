package main

import (
	"fmt"
	"math"
	"strings"
)

// Struct to represent a geometric shape
type Shape struct {
	Type string
}

// Circle is a struct representing a circle
type Circle struct {
	Shape
	Radius float64
}

// Rectangle is a struct representing a rectangle
type Rectangle struct {
	Shape
	Width  float64
	Height float64
}

// Function to calculate the area of a circle
func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Function to calculate the area of a rectangle
func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func main() {
	fmt.Println("Welcome to the Shape Area Calculator!")

	for {
		fmt.Print("Enter the shape (circle/rectangle) or 'quit' to exit: ")
		var input string
		fmt.Scanln(&input)

		// Convert input to lowercase for case-insensitive comparison
		input = strings.ToLower(input)

		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		} else if input == "circle" {
			fmt.Print("Enter the radius of the circle: ")
			var radius float64
			fmt.Scanln(&radius)
			circle := Circle{Shape: Shape{Type: "Circle"}, Radius: radius}
			area := circle.CalculateArea()
			fmt.Printf("Area of the circle: %.2f\n", area)
		} else if input == "rectangle" {
			fmt.Print("Enter the width of the rectangle: ")
			var width float64
			fmt.Scanln(&width)
			fmt.Print("Enter the height of the rectangle: ")
			var height float64
			fmt.Scanln(&height)
			rectangle := Rectangle{Shape: Shape{Type: "Rectangle"}, Width: width, Height: height}
			area := rectangle.CalculateArea()
			fmt.Printf("Area of the rectangle: %.2f\n", area)
		} else {
			fmt.Println("Invalid shape. Please enter 'circle' or 'rectangle'.")
		}
	}
}
