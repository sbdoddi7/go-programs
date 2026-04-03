package main

func main() {
	odd := make(chan int)
	even := make(chan int)
	go func() {
		for i := 1; i <= 10; i += 2 {
			odd <- i
		}
		close(odd)
	}()
	go func() {
		for i := 2; i <= 10; i += 2 {
			even <- i
		}
		close(even)
	}()
	for odd != nil || even != nil {
		select {
		case v, ok := <-odd:
			if ok {
				println("Odd:", v)
			} else {
				odd = nil
			}
		case v, ok := <-even:
			if ok {
				println("Even:", v)
			} else {
				even = nil
			}
		}
	}
}