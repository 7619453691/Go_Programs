// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x uint8 = 255
// 	x++ // overflow, x is 0
// 	fmt.Println(x)

// 	// a := int8(255 + 1)

// 	var b int8 = 127
// 	fmt.Printf("%d\n", b+1)

// 	b = -128
// 	b--
// 	fmt.Println(b)

// 	f := float32(math.MaxFloat32)
// 	fmt.Println(f)

// 	f = f * 1.2
// 	fmt.Println(f)

// 	//const i int8 = 300

// }

// *************** CONVERTING TYPES ***************************//

// package main

// import "fmt"

// func main() {
// 	var x = 3   // int type
// 	var y = 3.1 // float type

// 	//x = x * y

// 	x = x * int(y)
// 	fmt.Println(x)

// 	fmt.Printf("%T\n", int(y))

// 	x = int(float64(x) * y)
// 	fmt.Println(x)

// 	y = float64(x) * y
// 	fmt.Println(y)

// 	var a = 5 // int type
// 	fmt.Printf("%T\n", a)

// 	var b int64 = 2
// 	a = int(b)

// }

// *************8 CONVERTING NUMBER TO STRING AND STRING TO NUMBER **************//

package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := string(99)
	fmt.Println(s)

	// s1 := string(44.2)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	s1 := "3.123" // type string
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1 * 3.4)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)

}
