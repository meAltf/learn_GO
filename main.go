/** Learning Go languge **/

package main

import (
	"Learning/myutil"
	"fmt"
)

func main() {
	fmt.Println("Learning GO language | Yooo Robert!")
	fmt.Println("____________________________________")

	// Calling from another package
	myutil.PrintMessage("Hello from myutil directory..")
	myutil.JustPrint()
}
