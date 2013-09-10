/* 
Package pman is a password manager for the command line.

*/
package main

import (
	"fmt"
	"github.com/JonPulfer/pman/tools"
)

func main() {
	thisSecret := tools.HideInput("Password : ")
	fmt.Println(thisSecret)
}
