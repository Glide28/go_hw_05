package main

import (
	"github.com/Glide28/go_hw_05/gateway/internal/api"
	"github.com/Glide28/go_hw_05/ledger"
	"log"
	"net/http"
)

func main() {
	handler := api.NewHandler(ledger.New())
	log.Println("Gateway is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
