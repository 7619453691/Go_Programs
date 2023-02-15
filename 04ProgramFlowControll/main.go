// if else if, else statements are used for decission making

package main

import "fmt"

func main() {

	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive!")
	}

	//_ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// if price {
	// 	fmt.Println("something")
	// }

	if price < 100 {
		fmt.Println("It's cheep!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Its Expensive!")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years !\n", 18-age)
	} else if age == 18 {
		fmt.Printf("This is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote its important")
	} else {
		fmt.Printf("Invalid age!")
	}
}

// *******************************************///

// package main

// import "fmt"

// func main() {

//     // if condition_that_evaluates_to_boolean{
//     //      perform action1
//     // }else if condition_that_evaluates_to_boolean{
//     //      perform action2
//     // }else{
//     //      perform action3
//     // }

//     price, inStock := 100, true

//     if price >= 80 { // parenthesis are no required to enclose the testing condition
//         fmt.Println("Too Expensive")
//     }

//     if price <= 100 && inStock == true { //the same with: if price <= 100 && inStock { }
//         fmt.Println("Buy it!")
//     }

//     // In Go there is not such a thing like the Truthiness of a variable.
//     // Error:
//     // if price {
//     //  fmt.Println("We have price!")
//     // }

//     // only one if branch will be executed
//     if price < 100 {
//         fmt.Println("It's cheap!")
//     } else if price == 100 {
//         fmt.Println("On the edge")
//     } else { //executed only once if all the if branches are false (it's optional)
//         fmt.Println("It's Expensive!")
//     }
// }
