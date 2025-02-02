package main

import "fmt"
import "time"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)
	for i := 0; i < 10; i++ {
		go foo(fooVal, i)
	}
	for item := range fooVal {
		fmt.Println(item)
	}
	time.Sleep(time.Second * 2)
}
