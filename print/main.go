package main

import "fmt"

func main() {
	age := 25
	name := "Robert"
	height := 6.3467979

	fmt.Println("name: ", name, "age: ", age, "heigh: ", height)
	fmt.Println("Hey Robert!")

	// Printf | f-formatter | exact same thing like C language
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.4f\n", height)

	// to find TypeOf variables
	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of height is: %T\n", height)

	fmt.Printf("Name: %s, age: %d, height: %0.3f\n", name, age, height)

}
