package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"leopard.test/v2/internal/models"
	"leopard.test/v2/internal/repositories"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign models.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := repositories.CreateCampaign(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
        "success": true,
        "message": "Campaign created successfully",
        "data": id,
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
	   http.Error(w, "Invalid parameter", http.StatusInternalServerError)
	   return
	}

	campaign, err := repositories.GetCampaign(id)

	if err != nil {
		http.Error(w, "Invalid campaign", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(campaign)
}

func GetAllCampaigns(w http.ResponseWriter, r *http.Request) {

	campaigns, err := repositories.GetAllCampaigns()

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(campaigns)
}
