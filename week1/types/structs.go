package types

import "fmt"

func NewUser() {
	u := User{}

	//up 是指针
	up := &User{}
	up2 := new(User)
	//& 和 * 区别
	//&获取其指向操作数的指针
	//*表示指针指向的底层值

	fmt.Printf("u: %+v\n", u)
	fmt.Printf("up: %+v\n", up)
	fmt.Printf("up2: %+v \n", up2)

	u2 := User{Name: "Tom", Age: 16}
	fmt.Printf("u2: %+v \n", u2)

	u3 := &User{Name: "Jerry", Age: 8}
	fmt.Printf("u3: %+v \n", u3)
	u4 := *u3
	fmt.Printf("u3: %+v \n", u4)

}

type User struct {
	Name string
	Age  uint8
}
