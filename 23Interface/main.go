package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter:", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter:", r.perimeter())
// }

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	// c1 := circle{radius: 5}
	// r1 := rectangle{width: 3, height: 2.1}

	// printCircle(c1)
	// printRectangle(r1)

	// print(c1)
	// print(r1)

	// interface dynamic type and ploymorphism
	// var s shape
	// fmt.Printf("%T\n", s)

	// ball := circle{radius: 2.5}
	// s = ball

	// print(s)
	// fmt.Printf("Type of s: %T\n", s)

	// room := rectangle{width: 2, height: 3}
	// s = room
	// fmt.Printf("Type of s: %T\n", s)

	// TYPE ASSERTION AND TYPE SWITCHES
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	ball, ok := s.(circle)

	if ok == true {
		fmt.Printf("Ball volume: %v\n", ball.volume())
	}

	s = rectangle{width: 3.4, height: 2.2}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}

}

// ***************** code ****************************//

// package main

// import (
//     "fmt"
//     "math"
// )

// // declaring an interface type called shape
// // an interface contains only the signatures of the methods, but not their implementation
// type shape interface {
//     area() float64
//     perimeter() float64
// }

// // declaring 2 struct types that represent geometrical shapes: rectangle and circle

// type rectangle struct {
//     width, height float64
// }
// type circle struct {
//     radius float64
// }

// // method that calculates circle's area
// func (c circle) area() float64 {
//     return math.Pi * math.Pow(c.radius, 2)
// }

// // method that calculates rectangle's area
// func (r rectangle) area() float64 {
//     return r.height * r.width
// }

// // method that calculates circle's perimeter
// func (c circle) perimeter() float64 {
//     return 2 * math.Pi * c.radius
// }

// // method that calculates rectangle's perimeter
// func (r rectangle) perimeter() float64 {
//     return 2 * (r.height + r.width)
// }

// // any type that implements the interface is also of type of the interface
// // rectangle and circle values are also of type shape
// func print(s shape) {
//     fmt.Printf("Shape: %#v\n", s)
//     fmt.Printf("Area: %v\n", s.area())
//     fmt.Printf("Perimeter: %v\n", s.perimeter())
// }

// func main() {

//     // declaring a circle and a rectangle values
//     c1 := circle{radius: 5}
//     r1 := rectangle{width: 3, height: 2}

//     // circle and rectangle both implements the geometry interface  because they implement all methods of the interface
//     // an interface is implicitly implemented in Go
//     print(c1)
//     print(r1)

//     a := c1.area()
//     fmt.Println("Circle Area:", a)
// }
