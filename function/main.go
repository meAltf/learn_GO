package main

import "fmt"

func simpleFunction() {
	fmt.Println("DEMO simple function")
}

func addNum(a, b, c int) int {
	return a + b + c
}

func addNumWithDiffType(a int, b float64, c int) float64 {
	return float64(a) + b + float64(c)
}

// function with return
func multiply(a int, b int) (result int) {
	result = a * b
	return result
}

func main() {
	fmt.Println("Learning function in Go lamguage")
	simpleFunction()

	// return type function
	ans := addNum(3, 6, 9)
	fmt.Println("The sum:", ans)

	ans2 := addNumWithDiffType(2, 45.678, 78)
	fmt.Println("The new sum:", ans2)

	ansMult := multiply(5, 8)
	fmt.Println("The multiply:", ansMult)
}
