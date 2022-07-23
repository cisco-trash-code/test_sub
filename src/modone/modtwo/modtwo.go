package modtwo

import (
	"fmt"
	"modone/modthree"
)

func M2func() {
	fmt.Println("This is mod-two")
	modthree.SayHello()
}
