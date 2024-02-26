package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(pvk)

	puData := crypto.FromECDSAPub(&pvk.PublicKey)

	// crypto.

}
