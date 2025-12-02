package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"leopard.test/v2/internal/models"
	"leopard.test/v2/internal/repositories"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := repositories.CreateCustomer(customer)
	if err !=nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
        "success": true,
        "message": "Customer created successfully",
        "data": id,
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
	   http.Error(w, "Invalid parameter", http.StatusInternalServerError)
	   return
	}

	customer, err := repositories.GetCustomer(id)

	if err != nil {
		http.Error(w, "Customer Not Found", http.StatusNotFound)
		return 
	}
	
	json.NewEncoder(w).Encode(customer)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := repositories.GetCustomers()

	if err != nil {
		http.Error(w, "Error occurred", http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(customers)
}