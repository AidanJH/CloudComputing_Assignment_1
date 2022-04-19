package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateWallet() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//Strip the 0x from the front of the private key string
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	fmt.Println(publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if ok {
		fmt.Println(publicKeyECDSA)
	}

}
