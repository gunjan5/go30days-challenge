package main

import (
	"fmt"
)

type User struct {
	FName, LName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", i.FName, u.Lname)
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
}
