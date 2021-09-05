package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gravityblast/etherscan-go"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const etherscanKey = "YOUR_ETHERSCAN_API_KEY"

func executeLoop() {
	count := 0
	for {
		privateKey, err := crypto.GenerateKey()
		checkErr(err)
		privateKeyBytes := crypto.FromECDSA(privateKey)
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		if strings.HasPrefix(address, "0x0000") {
			if count%5 == 0 {
				time.Sleep(1 * time.Second)
			}
			c, err := etherscan.NewClient(etherscan.Mainnet, etherscanKey)
			checkErr(err)
			resp, err := c.Account(address)
			checkErr(err)
			balance := resp.Result
			fmt.Printf("Balance of wallet %v is %v\n", address, balance)
			if balance.Int64() != int64(0) {
				fmt.Printf("FOUND ETH IN WALLET: %v\n", address)
				fmt.Printf("Private Key: %v\n", hexutil.Encode(privateKeyBytes))
				break
			}
			count++
		}
	}
}

func main() {
	executeLoop()
}
