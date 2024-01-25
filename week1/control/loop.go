package control

import "fmt"

func ForLoop() {
	for i := 0; i < 10; {
		println(i)
		i++
	}
}

func WhileLoop() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func ForRange() {
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		println("index ", i, "value ", v)
	}
}

func MapLoop() {
	users := []User{
		{
			name: "Tome",
			age:  16,
		},
		{
			name: "Jerry",
			age:  14,
		},
	}
	m := make(map[string]*User, 2)
	for _, u := range users {
		m[u.name] = &u
	}

	for k, v := range m {
		fmt.Printf("name: %s, user %v", k, v)
	}
}

type User struct {
	name string
	age  int
}
