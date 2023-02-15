//01) 1. Using the var keyword declare a string called name and initialize it with your name.
//2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
//3. Print the following string on multiple lines like this:
//Your name: `here the value of name variable`
//Country: `here the value of country variable`
//4. Print out the following strings:
//a) He says: "Hello"
//b) C:\Users\a.txt

// package main

// import "fmt"

// func main() {
// 	var name = "PSBK"

// 	country := "India"

// 	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)

// 	// equivallent to
// 	fmt.Printf(`Your name: %s
// Country: %s
// `, name, country)

// 	fmt.Println("He says: \"Hello\"")
// 	fmt.Println("C:\\Users\\a.txt")
// }

//02) 1. Declare a rune called r that stores the non-ascii letter ă
//2. Print out the type of r
//3. Declare 2 strings that contains the values ma and m
//4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).
//BTW: mamă means mother in Romanian.
//Note: You should convert rune to string to contatenate it to another string.

// package main

// import "fmt"

// func main() {
// 	r := 'ă'
// 	fmt.Printf("r type: %T\n", r)

// 	s1, s2 := "ma", "m"

// 	s := s1 + s2 + string(r)
// 	fmt.Printf("s is %s\n", s) // s is  mamă
// }

//03) Consider the following string declaration: s1 := "țară means country in Romanian"
//1. Iterate over the string and print out byte by byte
//2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string

// package main

// import "fmt"

// func main() {
// 	s1 := "țară means country in Romanian"

// 	fmt.Printf("Bytes in string: ")
// 	for i := 0; i < len(s1); i++ {
// 		fmt.Printf("%v ", s1[i])
// 	}
// 	fmt.Println()

// 	for i, r := range s1 {
// 		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
// 	}
// 	fmt.Println()
// }

//04) find the errors

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	s1 := "Go is cool!"
// 	x := s1[0]
// 	fmt.Println(x)

// 	// ERROR -> cannot assign to s1[0]
// 	// s1[0] = 'x'

// 	// printing the number of runes in the string
// 	// fmt.Println(len(s1))
// 	fmt.Println(utf8.RuneCountInString(s1))
// }

//05) Consider the following string declaration:s := "你好 Go!"
//1. Convert the string to a byte slice.
//2. Print out the byte slice
//3. Iterate over the byte slice and print out each index and byte in the byte slice

// package main

// import "fmt"

// func main() {
// 	s := "你好 Go!"

// 	b := []byte(s)

// 	fmt.Printf("%#v\n", b)

// 	for i, v := range b {
// 		fmt.Printf("%d -> %d\n", i, v)
// 	}
// }

//06) Consider the following string declaration:s := "你好 Go!"
//1. Convert the string to a rune slice.
//2. Print out the rune slice
//3. Iterate over the rune slice and print out each index and rune in the rune slice

package main

import "fmt"

func main() {
	s := "你好 Go!"

	r := []rune(s)

	fmt.Printf("%#v\n", r)

	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
