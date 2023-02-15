// package main

// import "fmt"

// func main() {

// 	diana := struct {
// 		firstName, lastName string
// 		age                 int
// 	}{
// 		firstName: "Diana",
// 		lastName:  "Muller",
// 		age:       30,
// 	}

// 	fmt.Printf("%#v\n", diana)
// 	fmt.Printf("Dian's Age: %d\n", diana.age)

// 	type Book struct {
// 		string
// 		float64
// 		bool
// 	}

// 	b1 := Book{"1984 by george", 10.2, false}
// 	fmt.Printf("%#v\n", b1)

// 	fmt.Println(b1.string)

// 	type Employee struct {
// 		name   string
// 		salary int
// 		bool
// 	}

// 	e := Employee{"John", 40000, false}
// 	fmt.Printf("%#v\n", e)
// 	e.bool = true
// 	fmt.Printf("%#v\n", e)

// }

// ******** EMBEDED STRUCT ****************//

package main

import "fmt"

func main() {

	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "xyz",
		salary: 40000,
		contactInfo: Contact{
			email:   "xyz@gmail.com",
			address: "AHUHA SIUD 20, india",
			phone:   8765435678,
		},
	}
	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "aaa@gmail.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "andrei@gmail.com", phone: 7654356987, address: "India"}

	fmt.Println(myContact)

}
