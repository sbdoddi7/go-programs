package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// It is just declared and not initialized, so not pointing any underlying array
	var nilSlice []int

	// It is intialized with emplty array so it is null
	emptySlice := []int{}

	if nilSlice == nil {
		fmt.Println("yes it is nil slice")
	}

	if emptySlice != nil {
		fmt.Println("yes it is empty slice")
	}

	nilJsonSlice, _ := json.Marshal(nilSlice)

	emptyJsonSlice, _ := json.Marshal(emptySlice)

	fmt.Println("nil Json", string(nilJsonSlice))
	fmt.Println("empty Json", string(emptyJsonSlice))
}
