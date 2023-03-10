// package main

// import "fmt"

// func main() {
// 	name := "Andrei"
// 	fmt.Println(&name)

// 	var x int = 2
// 	ptr := &x

// 	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
// 	fmt.Printf("address of x is %p\n", &x)

// 	var ptr1 *float64
// 	_ = ptr1

// 	p := new(int)

// 	x = 100
// 	p = &x

// 	fmt.Printf("p is of tyoe %T with a value of %v\n", p, p)
// 	fmt.Printf("address of x is %p\n", &x)

// 	*p = 90 // equivallent to x = 90
// 	fmt.Println(x, *p)
// 	fmt.Println("*p == x:", *p == x)

// 	*p = 10     // x = 10
// 	*p = *p / 2 // x / 2
// 	fmt.Println(x)

// 	// &value ==> pointer
// 	// *pointer ==> value

// }

// ********** POINTER TO POINTER *****************//

// package main

// import "fmt"

// func main() {
// 	a := 5.5
// 	p1 := &a
// 	pp1 := &p1
// 	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
// 	fmt.Printf("Value of p1: %v, address of p1: %v\n", pp1, &pp1)

// 	fmt.Printf("*p1 is %v\n", *p1)
// 	fmt.Printf("*pp1 is %v\n", *pp1)

// 	fmt.Printf("**pp1 is %v\n", **pp1)

// 	**pp1++
// 	fmt.Printf("a is %v\n", a)

// 	// ***** comparing pointers ******//

// 	var p2 *int
// 	fmt.Printf("%#v\n", p2)

// 	y := 5
// 	p2 = &y

// 	z := 5
// 	p3 := &z

// 	fmt.Println(p2 == p3)

// 	p4 := &y
// 	fmt.Println(p2 == p4)

// }

