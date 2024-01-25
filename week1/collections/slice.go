package collections

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1: %v, len: %d, cap: %d, addr: %p \n", s1, len(s1), cap(s1), &s1)

	s2 := make([]int, 3, 4)
	fmt.Printf("s2: %v, len: %d, cap: %d, addr: %p \n", s2, len(s2), cap(s2), &s2)

	s2 = append(s2, 5)
	fmt.Printf("s2 appended: %v, len: %d, cap: %d, addr: %p \n", s2, len(s2), cap(s2), &s2)

	s2 = append(s2, 9)
	fmt.Printf("s2 appended: %v, len: %d, cap: %d, addr: %p \n", s2, len(s2), cap(s2), &s2)

}

func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len: %d, cap %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap %d \n", s2, len(s2), cap(s2))

	println("sub slice share same memory, change to sub slice also impact original slice")
	s2[2] = 99
	fmt.Printf("s1: %v, len: %d, cap %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap %d \n", s2, len(s2), cap(s2))

	println("if append to sub slice cause slice cap extended, the sub slice will be re-assign to new addr, no change to original slice anymore")
	s2 = append(s2, 50)
	s2[1] = 30
	fmt.Printf("s1: %v, len: %d, cap: %d, addr: %p \n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2: %v, len: %d, cap: %d, addr: %p \n", s2, len(s2), cap(s2), &s2)

}
