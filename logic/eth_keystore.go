package logic

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Keystore(){
//  key:=keystore.NewKeyStore("./wallet",keystore.StandardScryptN,keystore.StandardScryptP)
 password:="123"
//  acc, err:= key.NewAccount(password)
//  if err!=nil{
// 	log.Fatal(err)
//  }

//  fmt.Println(acc.Address)
file,err:= ioutil.ReadFile("./wallet/UTC--2023-07-30T18-18-36.947878346Z--4920ad0bb437ee99f3ea2a5e2d02d7f4d813fbc0")
if err!=nil{
	log.Fatal(err)
}
key,err2:= keystore.DecryptKey(file,password)

if err!=nil{
	log.Fatal(err2)
}

privateData:= crypto.FromECDSA(key.PrivateKey)
fmt.Println("private key",hexutil.Encode(privateData))

pubkeyData:= crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
fmt.Println("public key",hexutil.Encode(pubkeyData))

publicAddressData:= crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
fmt.Println("address of public key",publicAddressData)

}