package logic

import (
	"context"
	"fmt"
	"log"
	"math/big"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
    "encoding/hex"

)

var (
	url_goerli = "https://goerli.infura.io/v3/1b0d6b0a99704bbea5778b2de0938dd2"
	url_sepolia = "https://sepolia.infura.io/v3/1b0d6b0a99704bbea5778b2de0938dd2"
)

func GetEthers(){
	// keystore:=keystore.NewKeyStore("./wallet",keystore.StandardScryptN,keystore.StandardScryptP)
    // password:="123"
	// address1,err1:= keystore.NewAccount(password)
	// if err1!= nil{
	// 	log.Fatal(err1)
	// }
	// address2,err2:= keystore.NewAccount(password)
	// if err2!=nil{
    //    log.Fatal(err2)
	// }

	// fmt.Println("address 1",address1)
	// fmt.Println("address 2",address2)

	// 0xC58A42D58A1023329FEB7cDd66B612a2495f7E32 - address 1
	// 0xA93501be2Cf31993a53206fA36859B59F70733F7 - address 2


   client,err:=ethclient.Dial(url_sepolia)
   if err!=nil{
	log.Fatal(err)
   }
   defer client.Close()
   add1:= common.HexToAddress("0x10096e6F7c5f5B335abaA367016827e15DaFA65A")
   add2:= common.HexToAddress("0x7e21524B89b04B9f8e85591443f8B58379A969cF")

   balance1,_:=client.BalanceAt(context.Background(),add1,nil)
   balance2,_:=client.BalanceAt(context.Background(),add2,nil)
   
   nonce,_:= client.PendingNonceAt(context.Background(),add1) // this will give the nonce of the wallet from blockchain 
   fmt.Println("balance of wallet 1", balance1)
   fmt.Println("balance of wallet 2", balance2)

   // now we are sending tx from wallet 1 to wallet 2 
   // so we have to create a tranaction 
   // and to do that we will use this package called types 

   gasPrice,_:=client.SuggestGasPrice(context.Background())
   amount:= big.NewInt(100000000000000000)
   tx:= types.NewTransaction(nonce, add2,amount, 21000, gasPrice,nil)
    chainID,_:= client.NetworkID(context.Background())
  PrivateKey,_:= ParsePrivateKey("<your-privatekey>")
  
   transaction,_:=types.SignTx(tx,types.NewEIP155Signer(chainID),PrivateKey)
   client.SendTransaction(context.Background(),transaction)
  fmt.Println("tx hash",transaction.Hash().Hex())
}


func ParsePrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
    privateKeyBytes, err := hex.DecodeString(privateKeyHex)
    if err != nil {
        return nil, err
    }

    privateKey, err := crypto.ToECDSA(privateKeyBytes)
    if err != nil {
        return nil, err
    }

    return privateKey, nil
}






