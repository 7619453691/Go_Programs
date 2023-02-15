//01) 1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
//Print out the cities array and notice the value of its elements.
//2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
//Print out the grades array and notice the value of its elements.
//3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
//4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
//5. Iterate over grades using the range keyword and print out element by element.

// package main

// import "fmt"

// func main() {
// 	var cities [2]string
// 	fmt.Printf("%#v\n", cities)

// 	grades := [3]float64{12.3, 4.6}
// 	fmt.Printf("%#v\n", grades)

// 	taskDone := [...]bool{true, false, false, true}
// 	fmt.Printf("%#v\n", taskDone)

// 	for i := 0; i < len(cities); i++ {
// 		fmt.Println(cities[i])
// 	}

// 	for index, value := range grades {
// 		fmt.Println(index, value)
// 	}
// }

//02) Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}
//Write a Go program that counts the number of positive even numbers in the array.

// package main

// import "fmt"

// func main() {
// 	nums := [...]int{30, -1, -6, 90, -6}

// 	count := 0

// 	for _, v := range nums {
// 		if v%2 == 0 && v > 0 {
// 			count++
// 		}
// 	}
// 	fmt.Println("No. of positive even numbers in nums:", count)
// }

//03) Find errors

package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[0] = a
	myArray[0] = float64(a)

	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	//myArray[3] = 10.10
	myArray[2] = 10.10

	fmt.Println(myArray)

}
