//01) Create a new file in the current working directory called info.txt and then close the file. If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {

// 	newFile, err := os.Create("aa.txt")

// 	if err != nil {
// 		// log the error and exit the program
// 		log.Fatal(err) // the idiomatic way to handle errors
// 	}

// 	newFile.Close()
// }

//02) Rename the file created at Exercise #1 to data.txt
//Check if file exists before renaming it. If it doesn't exist print a message and stop the program.

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	// checking if file exists
// 	_, err := os.Stat("info.txt")
// 	// error handling
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			log.Fatal("The file does not exist")
// 		}
// 	}

// 	// renaming the file
// 	err = os.Rename("info.txt", "data.txt")

// 	// error handling
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//03) Remove the file created at exercise #1. Take care that the file is now called data.txt (it has been renamed at exercise #2).
//Perform error handling.

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	// removing the file
// 	err := os.Remove("data.txt")
// 	// error handling
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//05) Create a Go Program that reads the entire contents of a file called info.txt into a string. You can use ioutil.ReadAll() or any other function you want.
//The file exists in the current working directory.

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// )

// func main() {
// 	// opening the file in read-only mode. The file must exist (in the current working directory)
// 	// use a valid path!
// 	file, err := os.Open("info.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// defer closing the file
// 	defer file.Close()

// 	// ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read

// 	data, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Data as string: %s\n", data)
// }

//05) Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
//The file exists in the current working directory.

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	// opening the file in read-only mode. The file must exist (in the current working directory)
// 	// use a valid path!

// 	file, err := os.Open("info.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()

// 	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader

// 	scanner := bufio.NewScanner(file)

// 	// reading the whole file line by line:
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	// checking foe any possible error:
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// }

//06) Create a Go Program that:
//1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.
//2. Write the string "The Go gopher is an iconic mascot!" to the file.

// package main

// import (
// 	"io/ioutil"
// 	"log"
// )

// func main() {
// 	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
// 	// if the file doesn't exist WriteFile() creates it
// 	// and if it already exists the function will truncate it before writing to file.

// 	bs := []byte("The Go gopher is an iconic mascot!")
// 	err := ioutil.WriteFile("info.txt", bs, 0644)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
