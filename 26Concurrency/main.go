// // swapping go routines

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
	//(*wg).Done()
}

func f2() {
	fmt.Println("f2 execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	// fmt.Println("main execution started")
	// fmt.Println("No of cpu's:", runtime.NumCPU())
	// fmt.Println("No of goroutines:", runtime.NumGoroutine())

	// fmt.Println("OS:", runtime.GOOS)
	// fmt.Println("Arch:", runtime.GOARCH)

	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// go f1()
	// fmt.Println("No of goroutines after go f1():", runtime.NumGoroutine())

	// f2()

	// time.Sleep(time.Second * 2)
	// fmt.Println("main execution stopped")

	// *******WAIT GROUPS *************//
	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("No of goroutines after go f1():", runtime.NumGoroutine())

	f2()

	wg.Wait()
	fmt.Println("main execution stopped")

}
