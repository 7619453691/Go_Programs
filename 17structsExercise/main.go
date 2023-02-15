//01) 1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
//2. Declare and initialize two values of type person, one called me and another called you.
//3. Print out the struct values.

// package main

// import "fmt"

// type person struct {
// 	name   string
// 	age    int
// 	colors []string
// }

// func main() {

// 	me := person{name: "Maria", age: 30, colors: []string{"red", "yellow"}}

// 	you := person{name: "Gari", age: 22, colors: []string{"pink", "blue"}}

// 	fmt.Printf("%+v\n", me)
// 	fmt.Printf("%+v\n", you)

// }

//02) Consider the code from the previous exercise and:
//1. Change the name or the struct value called me to "Andrei"
//2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
//3. Add a new favorite color to the second struct value called you.
//4. Print out the struct values.

// package main

// import "fmt"

// type person struct {
// 	name   string
// 	age    int
// 	colors []string
// }

// func main() {

// 	me := person{name: "Maria", age: 30, colors: []string{"red", "yellow"}}

// 	you := person{name: "Gari", age: 22, colors: []string{"pink", "blue"}}

// 	me.name = "Andrei"

// 	var colors []string = you.colors
// 	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

// 	colors = append(colors, "green")
// 	you.colors = colors

// 	fmt.Printf("%+v\n", me)
// 	fmt.Printf("%+v\n", you)

// }

// 03) Consider the code from Exercise #1.
//Iterate and print out the favorite colors of the struct value called me.

// package main

// import "fmt"

// type person struct {
// 	name   string
// 	age    int
// 	colors []string
// }

// func main() {

// 	me := person{name: "Maria", age: 30, colors: []string{"red", "yellow"}}

// 	you := person{name: "Gari", age: 22, colors: []string{"pink", "blue"}}

// 	for index, color := range me.colors {
// 		fmt.Printf("%d -> %q\n", index, color)
// 	}

// 	_ = you
// }

//04) 1. Create a new struct type called grades with  2 fields: grade and course
//2. Add another field of type grades to person struct type (embedded struct).
//3. Change the initialization of the struct values called me and you to contain information about the grades.
//4. Change the grade and the course of one struct value to "Golang" and 98.
//5. Print out the struct values.

package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
	gr     grades
}

type grades struct {
	grade  int
	course string
}

func main() {

	me := person{
		name:   "Maria",
		age:    30,
		colors: []string{"red", "yellow"},
		gr:     grades{grade: 85, course: "Python"},
	}

	you := person{
		name:   "Gari",
		age:    22,
		colors: []string{"pink", "blue"},
		gr:     grades{grade: 100, course: "History"},
	}

	me.gr.grade = 98
	me.gr.course = "Golang"

	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)

}
