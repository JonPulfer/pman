package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleUserError(err error) {
	fmt.Printf("There has been an error: %v\n", err)
	os.Exit(1)
}

// Function AddKey gathers the details of the new key from the user and adds it to the key store
func AddKey(ks KeyStore, KeyID string, secret []byte) {
	var newKey Key
	ks.Open(secret)
	_, ok := ks[KeyID]
	if !ok {
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
	} else {
		fmt.Printf("Key already exists in the keystore. Use -e <keyname> to edit it\n")
	}
}

// Function EditKey guides the user through editing a key in the key store.
func EditKey(ks KeyStore, KeyID string, secret []byte) {
	var editKey Key
	ks.Open(secret)
	editKey, ok := ks[KeyID]

	if ok {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("What would you like to change?\n")
		fmt.Printf("Enter the first letter of each item you wish to edit: -\n")
		fmt.Printf("(L)oginname, (P)assword, (D)etail or (A)ll : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			handleUserError(err)
		}

		// Edit the Login name
		if strings.ContainsAny(input, "LA") {
			fmt.Printf("Enter the new login name: ")
			newLoginname, err := reader.ReadString('\n')
			if err != nil {
				handleUserError(err)
			}
			newLoginname = strings.Trim(newLoginname, "\n")
			editKey.LoginName = newLoginname
		}

		// Edit the password
		if strings.ContainsAny(input, "PA") {
			fmt.Printf("Enter the new password: ")
			newPassword, err := reader.ReadString('\n')
			if err != nil {
				handleUserError(err)
			}
			newPassword = strings.Trim(newPassword, "\n")
			editKey.OldPassword = editKey.Password
			editKey.Password = newPassword
		}

		// Edit the Detail
		if strings.ContainsAny(input, "DA") {
			fmt.Printf("Enter the new details: ")
			newDetail, err := reader.ReadString('\n')
			if err != nil {
				handleUserError(err)
			}
			newDetail = strings.Trim(newDetail, "\n")
			editKey.Detail = newDetail
		}

		ks[KeyID] = editKey

		ks.Close(secret)
	}
}
