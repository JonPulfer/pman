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
var addKey string

func init() {
	listKs = flag.Bool("l", false, "List the contents of the keystore")
	flag.StringVar(&addKey, "a", "", "Add a new key to the keystore")
	flag.Parse()
}

func main() {
	keyStore := make(tools.KeyStore, 1)
	thisSecret := tools.HideInput("Password : ")

	// Pad the key up to the expected block size
	key := []byte(thisSecret)
	for i := len(key); i < 32; i++ {
		key = append(key, []byte("a")[0])
	}
	fmt.Println()

	if *listKs {
		tools.ListKeystore(key)
	}

	if len(addKey) > 0 {
		tools.AddKey(keyStore, addKey, key)
	}

}
