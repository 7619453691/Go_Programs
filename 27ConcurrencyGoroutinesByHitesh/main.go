// GOroutines
//Thread : 1) Managed by os  2) fixed stack - 1mb

//Goroutines : 1) Managed by Go routine 2) Flexible stack - 2kb

// do not communicate by sharing memory, instead share memory by communicating

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
// func main() {

// 	go greeter("Hello")
// 	greeter("World")
// }

// ************* WAITGROUPS **************//

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer

var mut sync.Mutex // pointer

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d 200 status code for %s", res.StatusCode, endpoint)
	}
}

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}
