package main

import (
	"fmt"
)

func main() {
	var num1 int // the default value of veraible will be zero
	var num2 int
	num1 = 10
	num2 = 20
	var result = num1 + num2

	for i := 0; i <= 5; i++ {
		fmt.Println(result, ":", i)

	}

}
