package types

import "fmt"

func (u User) ChangeName(name string) {
	fmt.Printf("u in change name: %+v, %p \n", u, &u)
	u.Name = name
}

func (u *User) ChangeAge(age uint8) {
	fmt.Printf("u in change age: %+v, %p \n", u, u)
	u.Age = age
}

func ChangeUser() {
	fmt.Println(`Go中的参数传递是值传递：
	ChangeName中传递的是一份结构体的拷贝，接收器中改变的是拷贝的值，对原对象没影响
	ChangeAge传递的是指针，接收器能改变原对象的值`)
	u1 := User{Name: "Tom", Age: 18}
	fmt.Printf("u1 before change: %+v, %p \n", u1, &u1)
	u1.ChangeName("Jerry")
	u1.ChangeAge(20)
	fmt.Printf("u1 after change: %+v, %p \n", u1, &u1)
}
