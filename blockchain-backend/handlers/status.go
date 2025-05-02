package handlers

import (
	"blockchain-backend/models"
	"blockchain-backend/services"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
}

func StatusWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading WebSocket connection:", err)
		return
	}
	defer conn.Close()

	transactions := []*models.Transaction{
		{Hash: "0x123", From: "0xabc", To: "0xdef", Status: "loading"},
	}

	statusChannel := make(chan *models.Transaction)
	var wg sync.WaitGroup

	// Writer goroutine
	go func() {
		for tx := range statusChannel {
			err := conn.WriteJSON(tx)
			if err != nil {
				log.Println("Error sending WebSocket message:", err)
				return
			}
		}
	}()

	// Fetch statuses concurrently
	for _, txn := range transactions {
		wg.Add(1)
		go func(txn *models.Transaction) {
			defer wg.Done()
			status, err := services.GetTransactionStatus(txn.Hash)
			if err != nil {
				log.Printf("Error fetching status for tx %s: %v", txn.Hash, err)
				txn.Status = "Unknown"
			} else {
				txn.Status = status
			}
			statusChannel <- txn
		}(txn)
	}

	// Wait for all fetch goroutines to finish, then close channel
	go func() {
		wg.Wait()
		close(statusChannel)
	}()
}
