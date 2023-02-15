// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	// fmt.Println("os.Args", os.Args)

// 	// fmt.Println("Path:", os.Args[0])
// 	// fmt.Println("1st argument:", os.Args[1])
// 	// fmt.Println("2nd argument:", os.Args[2])
// 	// fmt.Println("No of items inside os.Args:", len(os.Args))

// 	var result, err = strconv.ParseFloat(os.Args[1], 64)
// 	fmt.Printf("%T\n", os.Args[1])
// 	fmt.Printf("%T\n", result)
// 	_ = err

// }

// ****************** SIMPLE IF STATEMENT ********************* ******************//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err = strconv.Atoi("20"); err == nil {
		fmt.Println("No error, i is:", i)
	} else {
		fmt.Println(err)
		//handle the error
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
		fmt.Printf("%T\n", args[1])
	}
}
