// package main

// import "fmt"

// func change(a *int) *float64 {
// 	*a = 100

// 	b := 5.5
// 	return &b
// }

// func changeVar(a int) {
// 	a = 66
// }

// func main() {

// 	x := 8
// 	p := &x

// 	fmt.Println("value of x before calling change():", x)
// 	change(p)
// 	fmt.Println("value of x after calling change():", x)

// 	fmt.Println("value of x before calling changeVar():", x)
// 	changeVar(x)
// 	fmt.Println("value of x after calling changeVar():", x)

// }

// ************* PASSING POINTERS TO FUNCTION PASSING BY VALUE VS PASSING BY FUNCTION *****************///

package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"

}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivallent to p.price = 300
	p.productName = "Bicycle"

}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValues():", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer():", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift)

	fmt.Println(gift)

	fmt.Println("Before calling changeProductByPointer:", gift)
	changeProductByPointer(&gift)
	fmt.Println("After calling changeProductByPointer:", gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println("price slice after calling changeSlice():", prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)
}
