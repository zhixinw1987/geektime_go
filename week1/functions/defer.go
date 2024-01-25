package functions

import "fmt"

//Defer is like finally in Java, execute the code block before exiting
//Defer define first execute last
func Defer() {
	println("open DS")
	defer func() {
		println("defer1: close DS")
	}()
	defer func() {
		println("defer2: close tranaction")
	}()
	println("start DB query")
	println("start transaction")
}

func DeferLoopV1() {
	println("loop v1")
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("val: %d addr: %p \n", i, &i)
		}()
	}
}

func DeferLoopV2() {
	println("loop v2")
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("val: %d addr: %p \n", val, &val)
		}(i)
	}
}

func DeferLoopV3() {
	println("loop v3")
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("val: %d addr: %p \n", j, &j)
		}()
	}
}
