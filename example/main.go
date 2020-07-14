package main

import (
	"GOFuture"
	"fmt"
	"sync"
	"time"
)

func testFuture(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Second)
	fmt.Println("fsdf")
}

func main() {

	future := GOFuture.Future{}
	future.RunAsync(testFuture)
	fmt.Println("asdfsdf")
	fmt.Println("asdfsdf")
	fmt.Println("asdfsdf")
	fmt.Println("asdfsdf")
	fmt.Println("asdfsdf")
	fmt.Println("asdfsdf")
	//future.Get()
	future.GetWithTimeOut(6 * time.Second)
}
