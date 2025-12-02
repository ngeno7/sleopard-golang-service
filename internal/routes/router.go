package routes

import (
    "github.com/gorilla/mux"
    "leopard.test/v2/internal/handlers"
)

func RegisterCustomerRoutes(router *mux.Router) {
    router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
    router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
}

func RegisterCampaignRoutes(router *mux.Router) {
    router.HandleFunc("/campaigns", handlers.GetAllCampaigns).Methods("GET")
    router.HandleFunc("/campaigns/{id}", handlers.GetCampaign).Methods("GET")
    router.HandleFunc("/campaigns", handlers.CreateCampaign).Methods("POST")
    router.HandleFunc("/campaigns/{id}/send", handlers.SendCampaign).Methods("POST")
}
