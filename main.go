package main

import "fmt"

type Customer struct {
	id   string
	Name string
	Age  int
}

type User struct {
	id   string
	Name string
	Age  int
}

func (c Customer) ID() string {
	return c.id
}

// func main() {
// 	var u = User{
// 		id:   "000",
// 		Name: "Jordy Polya",
// 		Age:  18,
// 	}

// 	var c = Customer(u)

// 	fmt.Printf("%T %#v", c, c)
// }

type Person interface {
	ID() string
}

type human func() string

func (h human) ID() string {
	return h()
}

func sayID(p Person) {
	fmt.Println(p.ID())
}

func main() {
	var fn = func() string {
		return "xxx"
	}

	//fmt.Println(human(fn).ID())
	sayID(human(fn))

	var fny = func() string {
		return "yyy"
	}
	//fmt.Println(human(fny).ID())
	sayID(human(fny))

}
