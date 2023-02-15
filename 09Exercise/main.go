//01) Using a composite literal declare and initialize a slice of type string with 3 elements.
//Iterate over the slice and print each element in the slice and its index.

// package main

// import "fmt"

// func main() {
// 	countries := []string{"India", "Brazil", "Germany"}

// 	for i, v := range countries {
// 		fmt.Printf("Index: %d, Element: %q\n", i, v)
// 	}

// }

//02) find errors

// package main

// import "fmt"

// func main() {
// 	mySlice := []float64{1.2, 5.6}

// 	// ERROR -> index out of range [2] with length 2
// 	// mySlice[2] = 6
// 	mySlice[1] = 6

// 	// ERROR -> cannot use a (type int) as type float64 in assignment
// 	// a := 10
// 	a := 10.
// 	mySlice[0] = a

// 	// ERROR -> index out of range [3] with length 2
// 	// mySlice[3] = 10.10

// 	mySlice[1] = 10.10

// 	mySlice = append(mySlice, a)

// 	fmt.Println(mySlice)

// }

//03) 1. Declare a slice called nums with three float64 numbers.
//2. Append the value 10.1 to the slice
//3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
//4. Print out the slice
//5. Declare a slice called n with two float64 values
//6. Append n to nums
//7. Print out the nums slice

// package main

// import "fmt"

// func main() {
// 	nums := []float64{1.1, 2.2, 3.3}

// 	nums = append(nums, 10.1)

// 	nums = append(nums, 4.1, 5.5, 6.6)

// 	fmt.Println(nums)

// 	n := []float64{10.10, 20.20}

// 	nums = append(nums, n...)

// 	fmt.Println(nums)

// }

//04) Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at the command line.
//The user should give between 2 and 10 numbers.
//Notes:
//- the program should be run in a terminal (go run main.go) not in Go Playground
//- example:
//go run main.go 3 2 5
//Expected output: Sum: 10, Product: 30

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func main() {

// 	if len(os.Args) < 2 {
// 		log.Fatal("Please run the program with arguments (2-10 numbers)!")
// 	}

// 	//taking the arguments in a new slice
// 	args := os.Args[1:]

// 	//declaring and initializing sum and product of type float64
// 	sum, product := 0., 1.

// 	if len(args) < 2 || len(args) > 10 {
// 		fmt.Println("Please enter between 2 and 10 numbers!")
// 	} else {
// 		for _, v := range args {
// 			num, err := strconv.ParseFloat(v, 64)
// 			if err != nil {
// 				continue // if it cant convert string to float64, continue with the next number
// 			}
// 			sum += num
// 			product *= num
// 		}
// 	}
// 	fmt.Printf("Sum: %v, Product: %v\n", sum, product)
// }

//05) Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
//Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
//Print those elements and their sum.

// package main

// import "fmt"

// func main() {
// 	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

// 	sum := 0

// 	for _, v := range nums[2 : len(nums)-2] {
// 		fmt.Println(v)
// 		sum += v
// 	}
// 	fmt.Println("Sum:", sum)
// }

//06) Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
//Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

// package main

// import "fmt"

// func main() {
// 	friends := []string{"Marry", "John", "Paul", "Diana"}

// 	yourFriends := make([]string, len(friends))

// 	copy(yourFriends, friends)

// 	yourFriends[0] = "Dan"
// 	fmt.Println(friends, yourFriends)
// }

//07) Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
//Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

// package main

// import "fmt"

// func main() {
// 	friends := []string{"Marry", "John", "Paul", "Diana"}

// 	yourFriends := []string{}

// 	yourFriends = append(yourFriends, friends...)

// 	yourFriends[0] = "Dan"
// 	fmt.Println(friends, yourFriends)
// }

//08) Consider the following slice declaration:
// years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
// Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}

// package main

// import "fmt"

// func main() {
// 	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
// 	newYears := []int{}
// 	newYears = append(years[:3], years[len(years)-3:]...)
// 	fmt.Printf("%#v\n", newYears)
// }
