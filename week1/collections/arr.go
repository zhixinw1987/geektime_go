package collections

import "fmt"

func Array() {
	arr1 := [3]int{7, 8, 9}
	fmt.Printf("a1: %v, len %d, cap %d \n", arr1, len(arr1), cap(arr1))

	arr2 := [3]int{5, 8}
	fmt.Printf("a2: %v, len %d, cap %d \n", arr2, len(arr2), cap(arr2))

	var arr3 [3]int
	fmt.Printf("a3: %v, len %d, cap %d \n", arr3, len(arr3), cap(arr3))
}
