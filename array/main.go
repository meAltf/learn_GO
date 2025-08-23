package main

import "fmt"

func main() {
	fmt.Println("Learning array in Go language")

	// 1. Array creation & initialization
	var name [5]string
	name[0] = "robert"
	name[1] = "broon"
	name[2] = "richard"

	fmt.Println("Names of persons is:", name)

	// 2. Array initialization dirrectly
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The numbers are:", numbers)

	fmt.Println("Length:", len(numbers))
	fmt.Println("value at index 2 in array:", numbers[2])

	// 3. default value of array is as similar to Java
	var names [5]string
	fmt.Printf("The default values are: %q\n", names)

	var booleanValues [6]bool
	fmt.Printf("The dafualt values for boolean: %v\n", booleanValues)
}
