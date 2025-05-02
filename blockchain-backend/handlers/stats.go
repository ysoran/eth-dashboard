package handlers

import (
	"encoding/json"
	"net/http"
)

type TopAddress struct {
	Address string `json:"address"`
	Value   int    `json:"value"`
}

type TransactionsPerBlock struct {
	BlockNumber int `json:"blockNumber"`
	TxCount     int `json:"txCount"`
}

type StatusRatio struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type ValueDistribution struct {
	Range string `json:"range"`
	Count int    `json:"count"`
}

type VolumeOverTime struct {
	BlockNumber int `json:"blockNumber"`
	TxCount     int `json:"txCount"`
}

// Handlers

func TopAddressesHandler(w http.ResponseWriter, r *http.Request) {
	data := []TopAddress{
		{"0xABC123...", 120},
		{"0xDEF456...", 95},
		{"0xGHI789...", 70},
	}
	json.NewEncoder(w).Encode(data)
}

func TxsPerBlockHandler(w http.ResponseWriter, r *http.Request) {
	data := []TransactionsPerBlock{
		{1001, 150},
		{1002, 170},
	}
	json.NewEncoder(w).Encode(data)
}

func StatusRatioHandler(w http.ResponseWriter, r *http.Request) {
	data := []StatusRatio{
		{"Success", 250},
		{"Failed", 45},
		{"Pending", 30},
	}
	json.NewEncoder(w).Encode(data)
}

func ValueDistributionHandler(w http.ResponseWriter, r *http.Request) {
	data := []ValueDistribution{
		{"0-1 ETH", 50},
		{"1-10 ETH", 120},
	}
	json.NewEncoder(w).Encode(data)
}

func VolumeOverTimeHandler(w http.ResponseWriter, r *http.Request) {
	data := []VolumeOverTime{
		{1001, 1000},
		{1002, 1050},
	}
	json.NewEncoder(w).Encode(data)
}
