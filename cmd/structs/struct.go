package main

import (
	"fmt"
	"time"
)

// use empty struct in method reciever when there is no need of data
type emptystruct struct {
}

func (e emptystruct) Hi() {
	fmt.Println("Calling Hi method empty struct")
}

func (e emptystruct) Hello() {
	fmt.Println("Calling Hello method empty struct")
}

func main() {
	// empty struct used in channel for signalling purpose
	// reason why we used the the empty struct is , containd no fields , no memory -> good for signalling purpose
	ch := make(chan struct{})

	go func() {
		fmt.Println("Sleeping for three seconds ...")
		time.Sleep(3 * time.Second)

		close(ch)
	}()

	<-ch

	fmt.Println("Tested channel behaviour with empty struct")

	var es emptystruct
	es.Hi()
	es.Hello()
}
