package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/**

		fmt.Println("What's your next Company?")
		var compName string

		fmt.Scan(&compName)
		fmt.Println("Hello Robert, your next could be:", compName)
	**/
	// To take input of a whole line
	fmt.Println("What's the list of your next Company?")
	reader := bufio.NewReader(os.Stdin)
	companyNameList, _ := reader.ReadString('\n')

	fmt.Println("Hello Robert, this is the list of companies:", companyNameList)
}
