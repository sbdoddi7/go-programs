package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		<-ch
	}()

	ch <- 10
	fmt.Println("done")
}
