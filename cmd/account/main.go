package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/algorand/go-algorand-sdk/crypto"
)

func Account() (account crypto.Account) {
	key := os.Getenv("PRIVATE_KEY")
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}

	account, err = crypto.AccountFromPrivateKey(decodedKey)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	fmt.Printf("Account is %s\n", Account().Address.String())
}
