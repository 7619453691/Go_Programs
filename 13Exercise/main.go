//01) 1. Declare a map called m1 which value is nil. Print out its type and value.
//2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.
//3. Add the following key: value to the map: 10: "Abba"
//4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.

// package main

// import "fmt"

// func main() {
// 	var m1 map[float64]bool
// 	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1)

// 	m2 := map[int]string{1: "Sting", 2: "Queen"}

// 	m2[10] = "Abba"

// 	fmt.Println(m2[2])   // existing key
// 	fmt.Println(m2[100]) // non existing key
// }

//02) find errors

// package main

// func main() {
// 	var m1 map[int]bool

// 	// ERROR -> panic: assignment to entry in nil map
// 	// m1[5] = true

// 	m2 := map[int]int{3: 10, 4: 40}
// 	m3 := map[int]int{3: 10, 4: 40}

// 	// ERROR -> invalid operation: m2 == m3 (map can only be compared to nil)
// 	// fmt.Println(m2 == m3)

// 	_, _, _ = m1, m2, m3
// }

//03) Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}
//1. The above map declaration has an error. Solve the error!
//2. Delete a key:value pair from the map.
//3. Iterate over the map and print out each key and value.

package main

import "fmt"

func main() {
	// m := map[int]bool{"1": true, 2: false, 3: false}
	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 2)

	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}

}
