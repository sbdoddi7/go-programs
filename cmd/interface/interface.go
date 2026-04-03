package main

import "fmt"

type empty_interface interface {
}

func explain(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Data is %T %d\n", v, v)
	case string:
		fmt.Printf("Data is %T  %s\n", v, v)
	case float64:
		fmt.Printf("Data is %T  %f\n", v, v)
	default:
		fmt.Println("Unsupported Data Type")
	}
}

func test_type_assertion(i any) {
	value, ok := i.(int)
	if !ok {
		fmt.Println("Data is different than expected")
		return
	}
	fmt.Printf("Data is %T %d\n", value, value)
}

func main() {
	// Zero value of empty interface is nil
	var ei empty_interface

	fmt.Println(ei)

	explain(10)
	explain("soma")
	explain(3.4)

	test_type_assertion("string")
}
