// package main

// import "fmt"

// func main() {

// 	numbers := []int{2, 3}

// 	numbers = append(numbers, 10)

// 	fmt.Println(numbers)

// 	numbers = append(numbers, 20, 30, 40)
// 	fmt.Println(numbers)

// 	n := []int{100, 200}

// 	numbers = append(numbers, n...)
// 	fmt.Println(numbers)

// 	src := []int{10, 20, 30}
// 	dst := make([]int, 2)
// 	nn := copy(dst, src)
// 	fmt.Println(src, dst, nn)

// }

// *********************** SLICE EXPRESSION ********************//

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a[1:3]

	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // same as s1[0:3]
	fmt.Println(s1[:])  // same as s1[0:len(s1)]

	// fmt.Println(s1[:100]) //slice bounds out of range [:100] with capacity 6

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)
}
