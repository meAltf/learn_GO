// Learning variables..
package myutil

import "fmt"

func Variables() {
	var name string = "Hey robert!"
	var version = 11
	fmt.Println(name)
	fmt.Println(version)

	var money int = 45600000
	var currency = 56
	fmt.Println(money)
	fmt.Println("currency: ", currency)

	var dimension float64 = 897.67
	fmt.Println(dimension)

	var decision bool = false
	fmt.Println(decision)

	var person = 23
	fmt.Println(person)

	const piValue = 67.34
	fmt.Println("pivalue: ", piValue)

	// popular way to create variables + shortcut
	dataVar := "shotcut way to create a variable!"
	fmt.Println(dataVar)

	// Important concept
	var Public = "starting with an uppercase letter is public (exported) — it can be used from another package."
	var private = "starting with a lowercase letter is private (unexported) — it cannot be accessed outside its package."

	fmt.Println(Public)
	fmt.Println(private)

}
