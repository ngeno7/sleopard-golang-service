package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"

    "leopard.test/v2/internal/routes"
)

func main() {
    router := mux.NewRouter()

    // link routes to handlers
    routes.RegisterCustomerRoutes(router)
	routes.RegisterCampaignRoutes(router)

    log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)
}