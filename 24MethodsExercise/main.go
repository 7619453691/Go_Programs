//01) 1. Create a new type called money. Its underlying type is float64.
//2. Create a method called print that prints out the money value with exact 2 decimal points.

// package main

// import "fmt"

// type money float64

// func (m money) print() {
// 	fmt.Printf("%.2f\n", m)
// }
// func main() {

// 	var salary money = 1.5 * 5.503
// 	fmt.Println(salary) // 8.2545

// 	salary.print()
// }

//02) Consider the money type declared at exercise #1. Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.

// package main

// import "fmt"

// type money float64

// func (m money) print() {
// 	fmt.Printf("%.2f\n", m)
// }

// func (m money) printStar() string {
// 	return fmt.Sprintf("%.2f", m)
// }

// func main() {
// 	var salary money = 1.5 * 5.503
// 	fmt.Println(salary)

// 	salary.print()

// 	s := salary.printStar()
// 	fmt.Println(s)
// }

//03) 1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
//2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
//3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.

// package main

// import "fmt"

// type book struct {
// 	title string
// 	price float64
// }

// func (b book) vat() float64 {
// 	return b.price * 0.09
// }

// func (b *book) discount() { //use a pointer receiver to change the struct value
// 	(*b).price = b.price * 0.9
// 	// equivallent to
// 	// b.price = b.price * 0.9
// }

// func main() {
// 	//book value
// 	bestBook := book{title: "The Trial", price: 9.9}

// 	// calling the method
// 	vat := bestBook.vat()
// 	fmt.Printf("Vat: %v\n", vat)

// 	bestBook.discount()
// 	fmt.Printf("%#v\n", bestBook)
// }

//04) A junior Go Programmer has just developed the following Go Program.   You want the change the book's price by calling changePrice method. However, you notice that after calling the method the price is not changed.
// You task is to change the code in order for the changePrice method to work as expected.

// package main

// import "fmt"

// type book struct {
//     title string
//     price float64
// }

// // method for book type
// func (b book) changePrice(p float64) {
//     b.price = p
// }

// func main() {
//     // book value
//     bestBook := book{title: "The Trial by Kafka", price: 9.9}

//     // changing the price
//     bestBook.changePrice(11.99)

//     fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
// }

// package main

// import "fmt"

// type book struct {
// 	title string
// 	price float64
// }

// // method for book type
// // the method should receive a pointer to book, not a value
// // in Go everyhing is passed by value so functions work on copies!
// func (b *book) changePrice(p float64) {
// 	b.price = p
// }

// func main() {
// 	// book value
// 	bestBook := book{title: "The Trial by Kafka", price: 9.9}

// 	// changing the price
// 	bestBook.changePrice(11.99)

// 	fmt.Printf("Book's price:%.2f\n", bestBook.price) // the price is changed
// }
