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
	
	k := tools.Key{LoginName: "pulfer", Password: "guff", Detail: "test entry"}
	ks := make(map[string]tools.Key)
	ks["test"]  = k
	tools.CreateStore(ks)
}
