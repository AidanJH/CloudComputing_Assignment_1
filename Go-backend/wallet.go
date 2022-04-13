package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/crypto/ecda"
)

func GenerateWallet() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

}
