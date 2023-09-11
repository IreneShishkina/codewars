package main

import "fmt"

func EvenOrOdd(number int) string {
	temp := number % 2
	if temp == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	num1 := 10
	res1 := EvenOrOdd(num1)
	fmt.Println(res1)
	num2 := 9
	res2 := EvenOrOdd(num2)
	fmt.Println(res2)
}
