package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./walley", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "pass@123"

	// _, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fileData, err := ioutil.ReadFile("./walley/UTC--2024-02-23T12-40-13.334660815Z--a2d87423bb1e961c741055c8038afdac52c76184")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(fileData, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub", hexutil.Encode(pData))

	fmt.Println("Add", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
