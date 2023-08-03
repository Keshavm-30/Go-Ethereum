package logic

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func EtherWallet(){
	pvk, err:= crypto.GenerateKey()
	if err!=nil{
		log.Fatal(err)
	}
fmt.Println(pvk)

	privateData:= crypto.FromECDSA(pvk)
	encoded_pvk:= hexutil.Encode(privateData)
	fmt.Println("private key generated",privateData)
	fmt.Println(encoded_pvk)

	//now we will generae a public key from this private key 
	publicKey:= crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(publicKey))

	//now we will generate the address from the public key 
	fmt.Println("address of the wallet ",crypto.PubkeyToAddress(pvk.PublicKey)) 

}