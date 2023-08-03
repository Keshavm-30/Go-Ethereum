package logic

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Generate_Key(){
	pvt_key,err:=crypto.GenerateKey()
	if err!=nil{
		log.Fatal(err)
	}

	// crypto.FromECDSA(pvt_key) // we will get a slice of byte 

	pvtKey_Data:= crypto.FromECDSA(pvt_key)
	fmt.Println("private key \n",hexutil.Encode(pvtKey_Data))
	

	publicKey_Data:= crypto.FromECDSAPub(&pvt_key.PublicKey)
	fmt.Println("public key of the generated private key \n",hexutil.Encode(publicKey_Data))

	publickey:=crypto.PubkeyToAddress(pvt_key.PublicKey).Hex()
	fmt.Println("address of the public key \n",publickey)
}