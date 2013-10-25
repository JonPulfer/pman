/*
Package pman is a password manager for the command line.

*/
package main

import (
	"flag"
	"fmt"
	"github.com/JonPulfer/pman/tools"
)

var listKs *bool

func init() {
	listKs = flag.Bool("l", false, "List the contents of the keystore")
	flag.Parse()
}

func main() {
	thisSecret := tools.HideInput("Password : ")
	fmt.Println()
	if *listKs {
		tools.ListKeystore(thisSecret)
	}
}
