// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}

// 	j := 10
// 	for j >= 0 {
// 		fmt.Println(j)
// 		j--
// 	}

// 	sum := 0
// 	for {
// 		sum++
// 	}
// 	fmt.Println(sum) // this line is never reached (infinite)

// }

// /////////////////////////*************** ////////////////////// //

// package main

// import "fmt"

// func main() {

//     // printing numbers from 0 to 9
//     for i := 0; i < 10; i++ {
//         fmt.Println(i)
//     }

//     // has the same effect as a while loop in other languages
//     // there is no while loop in Go
//     j := 10
//     for j >= 0 {
//         fmt.Println(j)
//         j--
//     }

//     // handling of multiple variables in a for loop
//     for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
//         fmt.Printf("i = %v, j = %v\n", i, j)
//     }

//     // infinite loop
//     // sum := 0
//     // for {
//     //  sum++
//     // }
//     // fmt.Println(sum) //this line is never reached
// }

// ********** FOR AND CONTINUE STATEMENT **************//

// package main

// import "fmt"

// func main() {

// 	for i := 0; i < 10; i++ {
// 		if i%2 != 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// **************** BREAK STATEMENT ***************///

package main

import "fmt"

func main() {
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Just a msg after the for loop")
}
