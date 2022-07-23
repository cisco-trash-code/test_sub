package main

import (
	"fmt"
	"modone/modtwo"
)

var x int

func main() {
	fmt.Println("This is main from mod-one")
	modtwo.M2func()
	fmt.Println("X on main:", x)
}
