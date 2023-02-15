//01) Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3)

// package main

// import "fmt"

// func cube(n float64) float64 {
// 	return n * n * n
// }
// func main() {

// 	x := cube(3)
// 	fmt.Println(x)
// }

//02) Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
//a) the factorial of n
//b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
//Test the program by calling the function.

// package main

// import "fmt"

// func f1(n uint) (int, int) {
// 	fact := 1
// 	sum := 0

// 	for i := 1; i <= int(n); i++ {
// 		fact *= i
// 		sum += i
// 	}
// 	return fact, sum
// }

// func main() {

// 	f, s := f1(4)
// 	fmt.Println("f:", f, "s:", s)

// 	f, s = f1(10)
// 	fmt.Println("f:", f, "s:", s)
// }

//03) Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nnn
//Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func myFunc(n string) int {
// 	// convert string to int
// 	i, err1 := strconv.Atoi(n)

// 	if err1 != nil {
// 		fmt.Printf("cannot convert %q to int.", n)
// 		os.Exit(1) // exit the program
// 	}

// 	ii, _ := strconv.Atoi(n + n)
// 	iii, _ := strconv.Atoi(n + n + n)

// 	return i + ii + iii
// }
// func main() {

// 	fmt.Println(myFunc("5"))
// }

//04) Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.

// package main

// import "fmt"

// func sum(a ...int) int {
// 	s := 0
// 	for _, v := range a {
// 		s += v
// 	}
// 	return s
// }

// func main() {
// 	s := sum(1, 2, 30)
// 	fmt.Println(s)
// }

//05) Change the function from the previous exercise and use a `naked return`.

// package main

// import "fmt"

// func sum(a ...int) (s int) {
// 	for _, v := range a {
// 		s += v
// 	}
// 	return
// }

// func main() {
// 	s := sum(1, 2, 30)
// 	fmt.Println(s)
// }

//06) Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.
// The function should search for the string (the second parameter) in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise. The function does a case-sensitive search.
//Call the function and see how it works.
//Example:
// animals := []string{"lion", "tiger", "bear"}
// result := searchItem(animals, "bear")
// fmt.Println(result) // => true
// result = searchItem(animals, "pig")
// fmt.Println(result)     // => false

// package main

// import "fmt"

// func searchItem(mySlice []string, myStr string) bool {
// 	for _, v := range mySlice {
// 		if v == myStr {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	animals := []string{"lion", "tiger", "bear"}
// 	result := searchItem(animals, "bear")
// 	fmt.Println(result) // => true
// 	result = searchItem(animals, "pig")
// 	fmt.Println(result) // => false
// }

//07) Change the function from the previous exercise to do a case-insensitive search.
//Example:
// animals := []string{"Lion", "tiger", "bear"}
// result := searchItem(animals, "beaR")
// fmt.Println(result) // => true
// result = searchItem(animals, "lion")
// fmt.Println(result)     // => true

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func searchItem(mySlice []string, myStr string) bool {
// 	for _, v := range mySlice {
// 		if strings.EqualFold(v, myStr) {
// 			return true
// 		}

// 	}
// 	return false
// }

// func main() {

// 	animals := []string{"lion", "tiger", "bear"}

// 	result := searchItem(animals, "BEAR")
// 	fmt.Println(result) // => true

// 	result = searchItem(animals, "lION")
// 	fmt.Println(result) // => true

// 	result = searchItem(animals, "pig")
// 	fmt.Println(result) // => false

// }

//08) package main

// import "fmt"

// func print(msg string) {
// 	fmt.Println(msg)
// }

// func main() {
// 	print("The Go gopher is the iconic mascot of the Go project.")
// 	fmt.Println("Hello, Go playground!")
// }
// Modify only the line in the main() body function where the print() function is invoked so that the program will print out Hello, Go playground! and then The Go gopher is the iconic mascot of the Go project.

// package main

// import "fmt"

// func print(msg string) {
// 	fmt.Println(msg)
// }

// func main() {
// 	defer print("The Go gopher is the iconic mascot of the Go project.")
// 	fmt.Println("Hello, Go playground!")
// }

//09) Create a function that takes in an int value and prints out that value.
//Assign the function to a variable, print out the type of the variable and then call that function through the variable name.

package main

import "fmt"

func f1(a int) {
	fmt.Println(a)
}

func main() {
	x := f1()
	fmt.Printf("%T\n", x) // => func(int)
	x(8)
}
