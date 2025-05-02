package models

// Transaction represents the Ethereum transaction structure.
type Transaction struct {
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Status      string `json:"status"`
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
}
