package main

import "fmt"

// creating function to understand Error handling and understand underscore identifier
func divideNum(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func divideNumSkipError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not ve zero")
	}
	return a / b, nil
}

func main() {

	fmt.Println("Learning error handling in function and understnd underscore identifies uses")
	result, err := divideNum(10, 0)

	if err != nil {
		fmt.Println("Error while divide - Need to handle")
	}
	fmt.Println("The final divide result:", result)

	// use underscore identifier to skip error or don't want to handle
	fmt.Println("____________________________________________________________________________________")
	result2, _ := divideNumSkipError(56, 0)
	fmt.Println("The final result after skipping the error:", result2)

}
