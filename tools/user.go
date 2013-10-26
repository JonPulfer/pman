package tools

import (
	"bufio"
	"fmt"
	"os"
)

func handleUserError(err error) {
	fmt.Printf("There has been an error: %v\n", err)
	os.Exit(1)
}

// Method AddKey gathers the details of the new key from the user and adds it to the key store
func AddKey(ks KeyStore, KeyID string, secret []byte) {
	var newKey Key
	ks.openKeystore(secret)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the username for this account: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		handleUserError(err)
	}
	newKey.LoginName = input
	fmt.Printf("and now if you could enter the password please: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		handleUserError(err)
	}
	newKey.Password = input
	fmt.Printf("finally, a bit of text to describe this account:\n")
	input, err = reader.ReadString('\n')
	if err != nil {
		handleUserError(err)
	}
	newKey.Detail = input

	ks[KeyID] = newKey
	ks.Close(secret)
}
