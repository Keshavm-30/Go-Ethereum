solc --bin --abi contract/basic.sol -o build // this command to create a build folder which contains basic.abi and basic.bin 

/home/keshav/go/bin/abigen --bin=build/basic.bin --abi=build/basic.abi 
--pkg=basic --out=gen/basic.go  // this command to genetare the golang code from smart contract 

whrere /home/keshav/go/bin/abigen is the  absolute path of abigen in my pc 