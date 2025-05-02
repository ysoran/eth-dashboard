package main

import (
	"blockchain-backend/handlers" // Correct import for handlers
	"blockchain-backend/utils"    // Correct import for utils
	"log"
	"net/http"
	"os"

	"github.com/rs/cors" // Import the CORS package
)

func main() {
	// Load environment variables from the .env file
	err := utils.LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	// Set up the routes and handlers
	http.HandleFunc("/api/transactions", handlers.TransactionsHandler) // Transaction endpoint
	http.HandleFunc("/ws/status", handlers.StatusWebSocketHandler)     // WebSocket handler for real-time updates
	http.HandleFunc("/api/stats/top-addresses", handlers.TopAddressesHandler)
	http.HandleFunc("/api/stats/txs-per-block", handlers.TxsPerBlockHandler)
	http.HandleFunc("/api/stats/status-ratio", handlers.StatusRatioHandler)
	http.HandleFunc("/api/stats/value-distribution", handlers.ValueDistributionHandler)
	http.HandleFunc("/api/stats/volume-over-time", handlers.VolumeOverTimeHandler)

	// Set up CORS with default options (allow all origins)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                      // Allow all origins
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Allow these methods
		AllowedHeaders: []string{"Content-Type"},           // Allow headers
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Start the server with CORS handling
	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(http.DefaultServeMux)))
}
