package tools

import (
	"fmt"
	"os"
)

// Type Key defines a single key containing login details for
// a particular account.
type Key struct {
	LoginName   string
	Password    string
	OldPassword string
	Detail      string
}

// Function CreateStore checks to make sure a store doesn't
// already exist before it creates a new one.
func CreateStore(k map[string]Key) {
	var kf string = "/home/jonathan/.pman/keystore"

	_, err := os.Stat(kf)
	if err != nil {
		kFile, err := os.Create(kf)
		if err != nil {
			panic(err)
		}
		defer kFile.Close()
		kFile.WriteString("test")
		fmt.Printf("\nCreated file\n")
	}

}
