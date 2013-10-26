package tools

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	kf string = "/home/jonathan/.pman/keystore"
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
		line = line + hdr + kn + x + k + e + ln + x + d.LoginName + e +
			pw + x + d.Password + e + op + x + d.OldPassword +
			e + dt + x + d.Detail + n
	}

	return line
}

// Function CreateStore checks to make sure a store doesn't
// already exist before it creates a new one.
func createStore(k KeyStore) {
	_, err := os.Stat(kf)
	if err != nil {
		kFile, err := os.Create(kf)
		if err != nil {
			panic(err)
		}
		defer kFile.Close()
		kFile.WriteString(k.String())
		fmt.Printf("\nCreated key store file\n")
	}

}

// Function List lists the contents of the keystore
func List(secret []byte) {
	ks := make(KeyStore)
	ks.Open(secret)
	for key, _ := range ks {
		fmt.Printf("Key: %s\n\tDetail: %s\n", key, ks[key].Detail)
	}
	ks.Close(secret)
}

// Function Query gets the details for 'qKey'.
func Query(qKey string, secret []byte) {
	ks := make(KeyStore)
	ks.Open(secret)
	fmt.Printf("Login:\t\t\t\t\tPassword:\t\tDetail:\n%s\t\t\t%s\t\t%s\n", ks[qKey].LoginName,
		ks[qKey].Password, ks[qKey].Detail)
}

// Method openKeystore opens the keystore
func (k *KeyStore) Open(secret []byte) {
	var newKeystore bool

	// Open the file containing the keystore

	kFile, err := os.Open(kf)
	if err != nil {
		createStore(*k)
		newKeystore = true

	}

	if !newKeystore {
		// Read the encrypted data from the file
		fInfo, err := kFile.Stat()
		if err != nil {
			panic("Cannot stat keystore file")
		}
		fileData := make([]byte, fInfo.Size())
		_, err = kFile.Read(fileData)
		if err != nil {
			panic(err)
		}
		kFile.Close()

		// Decrypt the data
		block, err := aes.NewCipher(secret)
		if err != nil {
			panic(err)
		}
		if len(fileData) < aes.BlockSize {
			panic("Encrypted data is not long enough to be a valid AES block")
		}
		iv := fileData[:aes.BlockSize]
		ciphertext := fileData[aes.BlockSize:]
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(ciphertext, ciphertext)

		// Decode the gob to return the KeyStore
		var buff bytes.Buffer
		_, err = buff.Write(ciphertext)
		if err != nil {
			panic(err)
		}
		dec := gob.NewDecoder(&buff)
		var kStore KeyStore
		err = dec.Decode(&kStore)
		if err != nil {
			fmt.Println("Data retrieved from store not valid, was the password correct?")
			os.Exit(1)
		}
		*k = kStore
	}
}

// Method Close encrypts and writes the encrypted data to the key file
func (k *KeyStore) Close(secret []byte) {
	// Encode a gob of the keystore
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(k)
	buffString := buff.String()

	// Encrypt the data
	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(buffString))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(buffString))

	// Write the encrypted data to the file
	kFile, err := os.OpenFile(kf, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	bytesWritten, err := kFile.Write(ciphertext)
	if err != nil || bytesWritten == 0 {
		panic(err)
	}
}

// Function Import loads records from a file to import into
// the keystore.
//
// It expects the records 1 per line in the following format
//
// keyname::loginname::password::oldpassword::detail
func Import(in string, secret []byte) {
	var k Key
	ks := make(KeyStore)
	ks.Open(secret)

	ifile, err := os.Open(in)
	if err != nil {
		fmt.Printf("Error opening import file : %s\n\t\t%v\n", in, err)
		return
	}
	defer ifile.Close()

	fscan := bufio.NewScanner(ifile)
	for fscan.Scan() {
		l := fscan.Text()
		v := strings.Split(l, "::")
		k.LoginName = v[1]
		k.Password = v[2]
		k.OldPassword = v[3]
		k.Detail = v[4]
		ks[v[0]] = k
	}
	ks.Close(secret)
}
