package services

import (
	"blockchain-backend/models" // Import the models package
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// GetLatestBlockNumber returns the latest Ethereum block number.
func GetLatestBlockNumber() (string, error) {
	client := resty.New()
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	// Make request to Infura to get the latest block number
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
			"jsonrpc": "2.0",
			"method": "eth_blockNumber",
			"params": [],
			"id": 1
		}`).
		Post(fmt.Sprintf("https://mainnet.infura.io/v3/%s", infuraProjectID))

	if err != nil {
		return "", err
	}

	var result struct {
		Result string `json:"result"`
	}
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return "", err
	}
	return result.Result, nil
}

// GetTransactionsForBlock fetches transactions for a specific Ethereum block number.
func GetTransactionsForBlock(blockNumber string) ([]*models.Transaction, error) {
	client := resty.New()
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	// Fetch the list of transactions for the specified block number
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "eth_getBlockByNumber",
			"params": ["%s", true],
			"id": 1
		}`, blockNumber)).
		Post(fmt.Sprintf("https://mainnet.infura.io/v3/%s", infuraProjectID))

	if err != nil {
		return nil, err
	}

	var result struct {
		Result struct {
			Transactions []map[string]interface{} `json:"transactions"`
		} `json:"result"`
	}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	// Map transactions into the custom Transaction struct
	var transactions []*models.Transaction
	for _, tx := range result.Result.Transactions {
		transaction := &models.Transaction{
			Hash:        tx["hash"].(string),
			From:        tx["from"].(string),
			To:          tx["to"].(string),
			Value:       tx["value"].(string),
			Status:      "loading", // Placeholder, update status after checking
			BlockHash:   tx["blockHash"].(string),
			BlockNumber: tx["blockNumber"].(string),
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetTransactionStatus fetches the status of a specific transaction.
func GetTransactionStatus(txHash string) (string, error) {
	client := resty.New()
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	// Request for detailed information about the transaction
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "eth_getTransactionByHash",
			"params": ["%s"],
			"id": 1
		}`, txHash)).
		Post(fmt.Sprintf("https://mainnet.infura.io/v3/%s", infuraProjectID))

	if err != nil {
		return "", err
	}

	var result struct {
		Result map[string]interface{} `json:"result"`
	}
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return "", err
	}

	if result.Result == nil {
		return "Not Found", nil
	}

	return "Success", nil
}
