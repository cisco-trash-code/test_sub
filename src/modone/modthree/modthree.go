package modthree

import "fmt"

var x int

func SayHello() {
	x = 15
	fmt.Println("Hello from mod-three")
	fmt.Println("X on mod-three:", x)
	Util3()
}
