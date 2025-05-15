package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// noResponse matches the JSON schema returned by the public NaaS HTTP API.
type noResponse struct {
	Reason string `json:"reason"`
}

// fetchNo performs a GET against the NaaS endpoint and returns the “no” reason.
func fetchNo() (string, error) {
	resp, err := http.Get("https://naas.isalman.dev/no")
	if err != nil {
		return "", fmt.Errorf("failed to fetch no: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("naas returned status %d", resp.StatusCode)
	}

	var nr noResponse
	if err := json.NewDecoder(resp.Body).Decode(&nr); err != nil {
		return "", fmt.Errorf("failed to decode no response: %w", err)
	}
	return nr.Reason, nil
}

func handleExcuse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		excuse := getRandomExcuse()
		json.NewEncoder(w).Encode(map[string]string{"excuse": excuse})

	case http.MethodPost:
		var body struct {
			Request string `json:"request"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		// If the incoming text contains "yes", get a hard no from NaaS.
		if strings.Contains(strings.ToLower(body.Request), "yes") {
			noMsg, err := fetchNo()
			if err != nil {
				log.Printf("error fetching no: %v", err)
				http.Error(w, "failed to get no", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"response": noMsg})
			return
		}

		// Otherwise, fall back to a random excuse.
		excuse := getRandomExcuse()
		json.NewEncoder(w).Encode(map[string]string{"excuse": excuse})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Seed the random number generator for getRandomExcuse.
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/excuse", handleExcuse)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	port := ":8080"
	log.Printf("Excuse-as-a-Service running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
