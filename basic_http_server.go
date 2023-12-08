package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultApiAddress = ":8080"
)

var (
	apiAddress string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	apiAddress = getVar("API_ADDRESS", defaultApiAddress)
}

func handlers() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/rate/btc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintf(w, "BitCoin to USD rate: %f $\n", 0.0); err != nil {
			log.Println(err)
		}
	})

	return r
}

func main() {
	log.Println(fmt.Sprintf("Server address: %s", apiAddress))
	log.Fatal(http.ListenAndServe(apiAddress, handlers()))
}

func getVar(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
