package deploy

import (
	"context"
	basic "ethereum/gen"
	"ethereum/logic"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url_sepolia = "https://sepolia.infura.io/v3/1b0d6b0a99704bbea5778b2de0938dd2"

func Deploy(){
client,err:=ethclient.Dial(url_sepolia)
if err!=nil{
	log.Fatal(err)
}

defer client.Close()

address:= common.HexToAddress("0x10096e6F7c5f5B335abaA367016827e15DaFA65A")
balance1,_:=client.BalanceAt(context.Background(),address,nil)
nonce,_:= client.PendingNonceAt(context.Background(),address)

fmt.Println("balance before the transaction",balance1)
gasPrice,_:=client.SuggestGasPrice(context.Background())

chainID,_:= client.NetworkID(context.Background())
PrivateKey,_:= logic.ParsePrivateKey("<your-private-key>")

auth,_:= bind.NewKeyedTransactorWithChainID(PrivateKey, chainID)

auth.GasPrice = gasPrice
auth.GasLimit= uint64(3000000)
auth.Nonce = big.NewInt(int64(nonce))
// SC_add,tx,_,_:= basic.DeployBasic(auth,client)

// fmt.Println("deployed smart contract address",SC_add.Hex())
// fmt.Println("transaction hash", tx.Hash().Hex())
balance2,_:=client.BalanceAt(context.Background(),address,nil)
fmt.Println("balance after the transaction",balance2)

smartcontractaddress:=common.HexToAddress("0x0B966dB007B834D29CccD37eBdD687BE4dFbFB6b")
sc_instance,_:=basic.NewBasic(smartcontractaddress,client)



tx,_:=sc_instance.Inc(auth)



fmt.Println(tx.Hash().Hex())
}

