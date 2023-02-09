package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Order struct {
	ID        int     `json:"id"`
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
}

type Customer struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age"`
	ContactInfo  ContactInfo `json:"contact_info"`
	Orders       []Order    `json:"orders"`
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var customers []Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process customer information
	for _, customer := range customers {
		log.Println("Processing customer:", customer)
		// Perform any additional processing here
	}

	w.WriteHeader(http.StatusOK)
	customerIDs := make([]int, len(customers))
	for i, customer := range customers {
		customerIDs[i] = customer.ID
	}
	response := fmt.Sprintf("Customers processed successfully: %v", customerIDs)
	w.Write([]byte(response))
}

func main() {
	log.Println("Starting HTTP server on port 8080")
	http.HandleFunc("/batch_customers", customerHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error:", err)
	}
}
 