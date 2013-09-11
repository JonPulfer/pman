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
	ks := make(tools.KeyStore)
	
	k1 := tools.Key{LoginName: "pulfer", Password: "guff", Detail: "test entry"}
	ks["test"] = k1
	k2 := tools.Key{LoginName: "pulfer@gmail.com", Password: thisSecret, Detail: "another fine test entry"}
	ks["test2"] = k2
	tools.CreateStore(ks)
}
