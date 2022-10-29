package main

import (
	"arteri-go-v1/controllers"
	"arteri-go-v1/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/records", controllers.GetRecords).Methods("GET")
	router.HandleFunc("/api/records/{id}", controllers.GetRecordById).Methods("GET")
	router.HandleFunc("/api/records", controllers.CreateRecord).Methods("POST")
	router.HandleFunc("/api/records/{id}", controllers.UpdateRecord).Methods("PUT")
	router.HandleFunc("/api/records/{id}", controllers.DeleteRecord).Methods("DELETE")

	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/{id}", controllers.GetCategoryById).Methods("GET")
	router.HandleFunc("/api/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/categories/{id}", controllers.DeleteCategory).Methods("DELETE")
}
