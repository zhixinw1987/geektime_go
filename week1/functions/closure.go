package functions

import "fmt"

func Closure() func() int {
	counter := 0
	fmt.Printf("init : %p", &counter)
	return func() int {
		counter++
		fmt.Printf("%p", &counter)
		println("")
		return counter
	}
}
