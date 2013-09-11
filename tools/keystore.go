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

type KeyStore map[string]Key

// Method String flattens the KeyStore into a single string
//
// The output format is:-
// :#{KEY}#:KeyName::somename:#:LoginName::username:#:Password::somepass:#:OldPassword::oldguff:#:Detail::This is the details:##:
func (ks *KeyStore) String() string {
	var line string

	for k, d := range *ks {
		hdr := ":#{KEY}#:"
		kn := "KeyName"
		ln := "LoginName"
		pw := "Password"
		op := "OldPassword"
		dt := "Detail"
		x := "::"
		e := ":#:"
		n := ":##:"
		line = line + hdr + kn + x + k + e + ln + x + d.LoginName + e + pw + x + d.Password +
			e + op + x + d.OldPassword + e + dt + x + d.Detail + n
	}

	return line
}

// Function CreateStore checks to make sure a store doesn't
// already exist before it creates a new one.
func CreateStore(k KeyStore) {
	var kf string = "/home/jonathan/.pman/keystore"

	_, err := os.Stat(kf)
	if err != nil {
		kFile, err := os.Create(kf)
		if err != nil {
			panic(err)
		}
		defer kFile.Close()
		kFile.WriteString(k.String())
		fmt.Printf("\nCreated file\n")
	}

}
