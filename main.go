package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StatusResponse struct {
	App       string `json:"app"`
	Category  string `json:"category"`
	Tech      string `json:"tech"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := StatusResponse{
			App:       "saas-billing-gateway-go-gin-v50",
			Category:  "SaaS Subscription Billing & Webhook Gateway",
			Tech:      "Go / Gin Engine",
			Timestamp: time.Now().Unix(),
		}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("[saas-billing-gateway-go-gin-v50] Go Microservice running on :8080...")
	http.ListenAndServe(":8080", nil)
}
