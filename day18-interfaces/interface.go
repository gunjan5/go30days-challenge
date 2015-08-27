package main

import (
	"fmt"
)

type User struct {
	FName, LName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FName, u.LName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}
func main() {
	u := &User{"Gunjan", "Patel"}
	fmt.Println(Greet(u))

	c := &Customer{12, "c3po"}
	fmt.Println(Greet(c))
}
