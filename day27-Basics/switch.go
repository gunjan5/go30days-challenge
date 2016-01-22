package main

import (
	"fmt"
)

func main() {

	name := "gunjan"

	switch {
	case len(name) == 6: //62
		fmt.Println("My name is Gunjan... Gunjan Patel")
		fallthrough
	case name == "notGunjan":
		fmt.Println("My name is not Gunjan")
	case name == "Gunjan":
		fmt.Println(`My name is Gunjan, with an uppercase "G"`) //representing the string as a rune type so it includes the ""s
	default:
		fmt.Println("bloody hell")

	}

	func_name(name)

	
}

func func_name(n Interface{}) {

	switch n.(type) {
	case int:
		fmt.Println("bloody int")
	case string:
		print("bloody string\n")
	case rune:
		print("bloody `rune`\n")
	default:
		fmt.Println("default")
	}
}
