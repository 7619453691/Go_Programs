// package main

// import "fmt"

// func main() {

// 	var numbers [4]int
// 	fmt.Printf("%v\n", numbers)
// 	fmt.Printf("%#v\n", numbers)

// 	var a1 = [4]float64{}
// 	fmt.Printf("%#v\n", a1)

// 	var a2 = [3]int{-10, 1, 100}
// 	fmt.Printf("%#v\n", a2)

// 	a3 := [4]string{"dan", "dayan", "paul", "john"}
// 	fmt.Printf("%#v\n", a3)

// 	a4 := [4]string{"x", "y"}
// 	fmt.Printf("%#v\n", a4)

// 	//ellipses operator
// 	a5 := [...]int{1, 2, 5, 1, -10, 66}
// 	fmt.Printf("%#v\n", a5)
// 	fmt.Printf("the length of a5 is %d\n", len(a5))

// 	a6 := [...]int{1,
// 		2,
// 		3,
// 		4,
// 		5, // this comma is mandatory
// 	}
// 	fmt.Printf("%#v\n", a6)

// }

///********************** ARRAY OPERATION ********************************//

package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// numbers[5] = 100 // error

	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)

}
