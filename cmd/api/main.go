package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"leopard.test/v2/internal/routes"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    router := mux.NewRouter()

    // link routes to handlers
    routes.RegisterCustomerRoutes(router)
	routes.RegisterCampaignRoutes(router)

    log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)

    // background to run in a different main script
    // BackgroundTask()
}