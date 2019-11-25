package main

import (
	"fmt"
)

func add(x, y int) (res1, res2 int) {
	res1 = x + y
	res2 = x - y
	return
}

func main() {
	var num1 int // the default value of veraible will be zero
	var num2 int
	num1 = 10
	num2 = 20
	var result = num1 + num2

	for i := 0; i <= 5; i++ {
		fmt.Println(result, ":", i)

	}

	myres1, myres2 := add(num1, num2)
	fmt.Println(myres1, myres2)

}
