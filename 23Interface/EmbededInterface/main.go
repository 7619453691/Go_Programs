// package main

// import (
// 	"fmt"
// 	"math"
// )

// type shape interface {
// 	area() float64
// }

// type object interface {
// 	volume() float64
// }

// // geometry is embeding shape and object interface
// type geometry interface {
// 	shape
// 	object
// 	getColor() string
// }

// type cube struct {
// 	edge  float64
// 	color string
// }

// func (c cube) area() float64 {
// 	return 6 * (c.edge * c.edge)
// }

// func (c cube) volume() float64 {
// 	return math.Pow(c.edge, 3)
// }

// func measure(g geometry) (float64, float64) {
// 	a := g.area()
// 	v := g.volume()
// 	return a, v
// }

// func (c cube) getColor() string {
// 	return c.color
// }

// func main() {

// 	c := cube{edge: 2}
// 	a, v := measure(c)
// 	fmt.Printf("Area: %v, Volume: %v\n", a, v)
// }

// ************* EMPTY INTERFACE *************************//

package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {

	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "go"
	fmt.Println(empty)

	empty = []int{1, 2, 4, 5}
	fmt.Println(empty)

	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "PSBK"
	you.info = 40
	you.info = []float64{1.1, 2.2}
	fmt.Println(you.info)
}
