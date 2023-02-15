//01)Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.

// package main

// import "fmt"

// func main() {

// 	for i := 1; i <= 50; i++ {
// 		if i%7 == 0 {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// 	fmt.Println("")
// }

//02) Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.

// package main

// import "fmt"

// func main() {

// 	for i := 1; i <= 50; i++ {
// 		if i%7 != 0 {
// 			continue
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println("")
// }

//03) Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.

// package main

// import "fmt"

// func main() {
// 	count := 0
// 	for i := 1; i <= 50; i++ {
// 		if i%7 != 0 {
// 			continue
// 		}
// 		fmt.Printf("%d ", i)
// 		count++

// 		if count == 3 {
// 			break
// 		}
// 	}
// 	fmt.Println("")
// }

//04) Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 500; i++ {
// 		if i%7 == 0 && i%5 == 0 {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// 	fmt.Println("")
// }

//05) Using a for loop print out all the years from your birthday to the current year.
//Use a variant of for loop where the post statement is moved inside the for block of code.

// package main

// import "fmt"

// func main() {

// 	birthYear, currentYear := 1999, 2022

// 	for i := birthYear; i <= currentYear; {
// 		fmt.Printf("%d ", i)
// 		i++
// 	}
// 	fmt.Println()
// }

//06) Change the code and use a switch statement instead of an if statement.

// package main
// import "fmt"

// func main() {
// 	age := -9
// 	if age < 0 || age > 100 {
// 		fmt.Println("Invalid Age")
// 	} else if age < 18 {
// 		fmt.Println("You are minor!")
// 	} else if age == 18 {
// 		fmt.Println("Congratulations! You've just become major!")
// 	} else {
// 		fmt.Println("You are major!")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	age := 10
// 	switch {
// 	case age < 0 || age > 100:
// 		fmt.Println("Invalid Age")
// 	case age <= 18:
// 		fmt.Println("You are minor!")
// 	case age == 18:
// 		fmt.Println("Congratulations! You've just become major!")
// 	default:
// 		fmt.Println("You are major!")
// 	}
// }
