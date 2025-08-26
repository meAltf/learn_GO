package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 21
	if x > 5 {
		fmt.Println("x is greate than")
	} else if x == 15 {
		fmt.Println("x is equal to 15")
	} else {
		fmt.Println("x is smaler than 5")
	}

	fmt.Println("____________________________________________________________________________________")
	fmt.Println("Take a inpput from user")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // remove newline

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid num, please enter a valid integer")
		return
	}

	if num > 100 {
		fmt.Println("Number is greater than 100")
	} else if num == 100 {
		fmt.Println("Number is eaxactly equal to 100")
	} else {
		fmt.Println("Number is less than 100")
	}

}
