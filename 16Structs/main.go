// package main

// import "fmt"

// func main() {

// 	title1, author1, year1 := "xyx", "Dinga", 1320
// 	title2, author2, year2 := "zzz", "Singa", 1606

// 	fmt.Println("Book1:", title1, author1, year1)

// 	fmt.Println("Book2:", title2, author2, year2)

// 	type book struct {
// 		title  string
// 		author string
// 		year   int
// 	}

// 	type book1 struct {
// 		title, author string
// 		year          int
// 	}

// 	myBook := book{"xyx", "Dinga", 1320}
// 	fmt.Println(myBook)

// 	bestBook := book{title: "Animal Farm", year: 1945, author: "George"}

// 	_ = bestBook

// 	aBook := book{title: "Random book"}
// 	fmt.Printf("%#v\n", aBook)
// }

// ********* RETRIVING AND UPDATING STRUCT **********///

package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "xyz"}
	fmt.Println(lastBook.title)

	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook) // + it will used to print field also

	aBook := book{title: "xyz", author: "Leo Tolstoy", year: 1878}

	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)

}
