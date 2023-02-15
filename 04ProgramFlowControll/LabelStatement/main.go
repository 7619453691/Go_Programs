// package main

// import "fmt"

// func main() {

// 	// labels dont conflicts
// 	// outer := 19
// 	// _ = outer
// 	people := [5]string{"Raki", "Mark", "Dingi", "Singi", "Gaja"}
// 	friends := [2]string{"Dingi", "Marry"}

// outer:
// 	for index, name := range people {
// 		for _, friend := range friends {
// 			if name == friend {
// 				fmt.Printf("Found a friend %q at index %d\n", friend, index)
// 				break outer
// 			}
// 		}
// 	}
// 	fmt.Println("Next instruction after the break")
// }

/// ****************GOTO STATEMENT ********************///

package main

import "fmt"

func main() {

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// This is an error
	// 	goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("something here!")
}
