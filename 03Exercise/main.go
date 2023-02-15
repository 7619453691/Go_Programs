//01) Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
//Using short declaration syntax declare and assign these values to variables x, y and z:
//- 20
//- 15.5
//- "Gopher!"
//Using fmt.Println() print out the values of a, b, c, d, x, y and z.
//Try to run the program without error.

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b float64
// 	var c bool
// 	var d string

// 	x := 20
// 	y := 15.5
// 	z := "Gopher"

// 	fmt.Println(a, b, c, d)
// 	fmt.Println("x is:", x, "y is:", y, "z is:", z)
// }

//02Change the code from the previous exercise in the following way:
//1. Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.
//2. Declare x, y and z on a single line -> multiple short declarations
//3. Remove the statement that prints out the variables. See the error!
//4. Change the program to run without error using the blank identifier (_)

// package main

// func main() {

// 	var (
// 		a int
// 		b float64
// 		c bool
// 		d string
// 	)

// 	x, y, z := 20, 15.5, "Gophers"
// 	_, _, _, _, _, _, _ = a, b, c, d, x, y, z
// }

//03) find errors

// package main

// func main() {
// 	var a float64 = 7.1

// 	x, y := true, 3.7

// 	// ERROR -> no new variables on left side of :=
// 	// a, x := 5.5, false

// 	a, x = 5.5, false
// 	_, _, _ = a, x, y
// }

// find errors

// package main

// import "fmt"

// // ERROR -> := is not allowed in package scope (only local scope)
// // version := "3.1"

// func main() {
// 	// ERROR -> A string is initialized only using double quotes ("")
// 	// name := 'Golang'

// 	name := "Golang"
// 	fmt.Println(name)
// }
