package main

import (
	"fmt"
)

func demo() (a int) {
	a = 9
	fmt.Println(a)
	return
}

func main() {
	sam := demo()
	fmt.Println(sam)
}
