package main

import (
	"fmt"
	"net/http"
	"os"
	"distributed-log-monitoring/backend/internal/logger"
	"distributed-log-monitoring/backend/internal/config"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()
    // Initialize logger
	log, err := logger.Init(cfg.LogLevel, cfg.LogFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	// Initialize metrics
	_, err = metrics.Init()
	if err != nil {
		log.Fatalf("Failed to initialize metrics: %v", err)
	}

	log.Infof("Starting %s v%s", cfg.AppName, cfg.AppVersion)
	log.Debugf("Environment: %s, Debug: %v", cfg.Environment, cfg.Debug)





	// 1. Read configuration port from environment variables (defaults to 8000)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}

	// 2. Add the /health endpoint required by your docker-compose healthcheck
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// 3. Add a base test route for your frontend API calls
	http.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Distributed Log Monitoring API is live"}`))
	})

	// 4. Start the server loop to keep the container running indefinitely
	fmt.Printf("Backend server listening on 0.0.0.0:%s...\n", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
