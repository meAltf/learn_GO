package main

import "fmt"

func main() {
	fmt.Println("Learing slice in Go language | same like arrayList in java")
	fmt.Println("____________________________________________________________________________________")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are:", numbers)
	fmt.Printf("Numbers has data ttype of: %T\n", numbers)
	fmt.Println("The length of numbers array:", len(numbers))

	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 99)
	fmt.Println("After append the final numbers array:", numbers)
	fmt.Println("After append, the length of numbers array:", len(numbers))

	// 1. Initialize empty slice | means with 0 data
	stock := []int{}
	fmt.Println("slice:", stock)
	fmt.Println("length", len(stock))
	fmt.Println("capacity", cap(stock))

	// 2. Initialize empty slice | means with 0 data, 0 capacity
	stockMake := make([]int, 0)
	fmt.Println("slice:", stockMake)
	fmt.Println("length", len(stockMake))
	fmt.Println("capacity", cap(stockMake))

	// 3. Initialize a slice with some data and capacity
	fmt.Println("____________________________________________________________________________________")
	money := make([]int, 3, 5)
	fmt.Println("slice:", money)
	fmt.Println("length", len(money))
	fmt.Println("capacity", cap(money))

	money = append(money, 4, 5)
	fmt.Println("slice:", money)
	fmt.Println("length", len(money))
	fmt.Println("capacity", cap(money))

	// Now, it will increase capacity = 2*currentCapacity | bcz size exceeds more than 5(current capacity)
	money = append(money, 99)
	fmt.Println("slice:", money)
	fmt.Println("length", len(money))
	fmt.Println("capacity", cap(money))
}
