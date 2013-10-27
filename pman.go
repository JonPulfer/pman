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
var editKey string
var addKey string
var delKey string
var iFile string
var qKey string

func init() {
	listKs = flag.Bool("l", false, "List the contents of the keystore")
	flag.StringVar(&editKey, "e", "", "Edit <KeyID> entry")
	flag.StringVar(&addKey, "a", "", "Add a new key to the keystore")
	flag.StringVar(&delKey, "d", "", "Delete <KeyID> from the keystore")
	flag.StringVar(&iFile, "i", "", "Import entries from file")
	flag.StringVar(&qKey, "q", "", "Query the store for key")
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
		tools.List(key)
	}

	if len(addKey) > 0 {
		tools.AddKey(keyStore, addKey, key)
	}

	if len(iFile) > 0 {
		tools.Import(iFile, key)
	}

	if len(qKey) > 0 {
		tools.Query(qKey, key)
	}

	if len(editKey) > 0 {
		tools.EditKey(keyStore, editKey, key)
	}

	if len(delKey) > 0 {
		tools.Delete(delKey, key)
	}

}
