// package main

// import (
// 	"fmt"
// )

// func main() {
// 	const days int = 7

// 	var i int
// 	fmt.Println(i)

// 	const pi float64 = 3.14
// 	const secondInHour = 3600

// 	duration := 234 // in hours
// 	fmt.Printf("Duration in seconds: %v\n", duration*secondInHour)

// 	x, y := 5, 1
// 	fmt.Println(x / y)

// 	const a = 5
// 	const b = 0
// 	// fmt.Println(a / b) // error

// 	const n, m int = 4, 5
// 	const n1, m1 = 6, 7

// 	const (
// 		min1 = -500
// 		min2
// 		min3
// 	)
// 	fmt.Println(min1, min2, min3)

// 	//Constant Rules
// 	//1. You cannot change a constant
// 	const temp int = 100
// 	// temp = 50

// 	//2.You cannot initiate a constant at runtime
// 	// const power = math.Pow(2, 3)

// 	//3.You cannot use a variable to initialize a constant
// 	// t := 5
// 	// const tc = t

// 	//4.
// 	const l1 = len("Hello")

// }

//************** Type and untype constant ***********************

// package main

// import "fmt"

// func main() {
// 	const a float64 = 5.1 // typed constant
// 	const b = 6.7         // untyped constant

// 	const c float64 = a * b
// 	const str = "Hello " + "Go!"

// 	const d = 5 > 10
// 	fmt.Println(d)

// 	// const x int = 5
// 	// const y float64 = 2.2 * x

// 	const x = 5
// 	const y = 2.2 * 5
// 	fmt.Printf("%T\n", y)

// 	var i int = x     // x changes to int
// 	var j float64 = x // var j float64 = float64(x)
// 	var p byte = x

// 	fmt.Println(i, j, p)

// 	const r = 5
// 	var rr = r
// 	fmt.Printf("%T\n", rr)
// }

// ***************** IOTA **********************************//

// package main

// import "fmt"

// func main() {
// 	const (
// 		c1 = iota
// 		c2 = iota
// 		c3 = iota
// 	)

// 	fmt.Println(c1, c2, c3)

// 	const (
// 		c11 = iota
// 		c22
// 		c33
// 	)
// 	fmt.Println(c11, c22, c33)

// 	const (
// 		North = iota // default 0
// 		East
// 		South
// 		West
// 	)
// 	fmt.Println(West)

// 	const (
// 		a = (iota * 2) + 1 // 1
// 		b                  // 3
// 		c                  // 5
// 	)
// 	fmt.Println(a, b, c)

// 	// x = -2, y = -4, z = -5
// 	const (
// 		x = -(iota + 2) // -2
// 		_               //-3
// 		y               //-4
// 		z               //-5
// 	)
// 	fmt.Println(x, y, z)
// }
