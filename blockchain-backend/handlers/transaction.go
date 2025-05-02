// handlers/transactions.go
package handlers

import (
	"blockchain-backend/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TransactionsHandler hit")
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Fetch the latest block number from Ethereum
	blockNumber, err := services.GetLatestBlockNumber()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch latest block number: %v", err), http.StatusInternalServerError)
		log.Printf("Error fetching latest block number: %v", err)
		return
	}

	// Fetch transactions for the latest block
	transactions, err := services.GetTransactionsForBlock(blockNumber)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch transactions for block %s: %v", blockNumber, err), http.StatusInternalServerError)
		log.Printf("Error fetching transactions for block %s: %v", blockNumber, err)
		return
	}

	if transactions == nil {
		log.Println("No transactions found for block", blockNumber)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]string{"No transactions found for the latest block"})
		if err != nil {
			log.Printf("Error encoding response: %v", err)
		}
		return
	}

	// Log the transactions to verify if they are fetched properly
	log.Printf("Fetched %d transactions for block %s", len(transactions), blockNumber)

	// Respond with the transactions in JSON format immediately (showing loading status)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Printf("Error encoding transactions to JSON: %v", err)
	}
}
