package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Card struct {
	Number string `json:"number"`
}

type ValidationResult struct {
	Valid bool `json:"valid"`
}

func isValidCreditCard(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	var sum int
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if double {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var card Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
	}

	result := ValidationResult{
		Valid: isValidCreditCard(card.Number),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/validate", validateCardHandler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
