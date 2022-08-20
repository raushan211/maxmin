package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("enter nos.")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println("maximum", max(num1, num2))
	fmt.Println("minimum", min(num1, num2))

}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func min(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
