//01) Consider the following variable declaration x := 10.10
//1. Print out the address of x
//2. Declare a pointer called ptr that stores the address of x.
//3. Print out the type and the value of ptr.
//4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).

// package main

// import "fmt"

// func main() {

// 	x := 10.10
// 	fmt.Printf("The address of x is %p\n:", &x)

// 	ptr := &x

// 	fmt.Printf("ptr type: %T, ptr value: %v\n", ptr, ptr)

// 	fmt.Printf("The address of ptr: %p\n", &ptr)
// 	fmt.Printf("The value of x through the pointer: %f\n", *ptr)

// }

//02) Consider the following variable declarations:
//x, y := 10, 2
//ptrx, ptry := &x, &y
//Declare a new variable called z and initialize the variable by dividing x by y through the pointers.

// package main

// import "fmt"

// func main() {
// 	x, y := 10, 2
// 	ptrx, ptry := &x, &y

// 	z := *ptrx / *ptry
// 	fmt.Printf("z is %d\n", z)
// }

//03) Consider the following variable declaration:x, y := 5.5, 8.8
//Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5

package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}
func main() {
	x, y := 5.5, 8.8
	swap(&x, &y)

	fmt.Printf("x is %v and y is %v\n", x, y)

}
