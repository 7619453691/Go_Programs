// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func foo() {
// 	fmt.Println("This is foo()")
// }

// func bar() {
// 	fmt.Println("This is bar()")
// }

// func foobar() {
// 	fmt.Println("This is foobar()")
// }

// func main() {
// 	defer foo()
// 	bar()

// 	fmt.Println("Just a string after defering foo() and calling bar()")

// 	defer foobar()

// 	file, err := os.Open("main.go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()

// 	// code that works (read/write) with the file
// }

// *********************.// **********************

// package main

// import (
//     "fmt"
// )

// func foo() {
//     fmt.Println("This is foo()!")
// }

// func bar() {
//     fmt.Println("This is bar()!")
// }

// func foobar() {
//     fmt.Println("This is foobar()!")
// }

// func main() {

//     // a defer statement defers or postpones the execution of a function until the surrounding function returns.

//     // by deferring foo() it will execute it just before exiting the surrounding function which is main()
//     defer foo()
//     bar()

//     fmt.Println("Just a string after deferring foo() and calling bar()")

//     // if there are more functions deferred, Go will execute them in the reverse order they were deferred
//     defer foobar()

//     /*
//         When executing the program the fallowing output is printed out:

//         This is bar()!
//         Just a string after deferring foo() and calling bar()
//         This is foobar()!
//         This is foo()!
//     */

// }

// ************ ANONYMOUS FUNCTION **********************//

package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	func(msg string) {
		fmt.Println(msg)
	}("i'm an anonymous function!")

	a := increment(10)
	fmt.Printf("%T\n", a)

	a()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}
