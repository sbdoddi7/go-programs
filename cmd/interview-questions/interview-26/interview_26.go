package main

import "fmt"

type X struct {
	data []int
}

func modify(x X) {
	x.data[0] = 100
}

func main() {
	obj := X{data: []int{1, 2, 3}}
	modify(obj)
	fmt.Println(obj.data[0])
}