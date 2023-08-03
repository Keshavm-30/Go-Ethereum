package logic
import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "math/big"
)

var url_maiinet = "https://mainnet.infura.io/v3/1b0d6b0a99704bbea5778b2de0938dd2"
var ganache_url = "http://localhost:8545"

func Get_Info(){
	var wallet_addr = "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"
	address:= common.HexToAddress(wallet_addr)


	client,err:=ethclient.DialContext(context.Background(),url_maiinet)// ethclient allows us to create new ethereum client 

	if err!=nil{
		log.Fatal("error to create a new client",err)
    }
	defer client.Close()
	// blockNumber := big.NewInt(8976)
	block,err:=client.BlockByNumber(context.Background(),nil) // nil means it will return the last blocknumber 
	if err!=nil{
		log.Fatal("error loding the latest blocknumber",err)
	}
	fmt.Println("current block in ethereum",block.Number())
	fmt.Println("transaction hash of current block in ethereum",block.TxHash())

	// fmt.Println(block.Body().Transactions)

	balance,err:=client.BalanceAt(context.Background(),address,nil)
	if err!=nil{
		log.Fatal("unable to fethch the balance of given addrss",err)
	}
	fmt.Println("balance of wallet is", balance)

}
