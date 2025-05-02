package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
)

func GetTotalTransactions() int {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	return len(block.Transactions())
}

func GetTotalValue() float64 {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	var totalValue float64
	for _, tx := range block.Transactions() {
		totalValue += float64(tx.Value().Uint64())
	}

	return totalValue / 1e18 // Convert to Ether
}

// GetPendingTransactions fetches the number of pending transactions via Infura's RPC endpoint.
func GetPendingTransactions() int {
	client := resty.New()
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	// Request the list of pending transactions (use eth_pendingTransactions)
	body := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_pendingTransactions",
		"params": [],
		"id": 1
	}`)

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(fmt.Sprintf("https://mainnet.infura.io/v3/%s", infuraProjectID))

	if err != nil {
		log.Fatal(err)
	}

	var result struct {
		Result []map[string]interface{} `json:"result"`
	}
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		log.Fatal(err)
	}

	return len(result.Result)
}

// GetTransactionsOverTime fetches transaction counts over the last 1000 blocks.
func GetTransactionsOverTime() []map[string]interface{} {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var transactionsOverTime []map[string]interface{}
	for i := blockNumber - 1000; i <= blockNumber; i++ {
		block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		if err != nil {
			log.Fatal(err)
		}
		transactionsOverTime = append(transactionsOverTime, map[string]interface{}{
			"blockNumber":  i,
			"transactions": len(block.Transactions()),
		})
	}

	return transactionsOverTime
}
