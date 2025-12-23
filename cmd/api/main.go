package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utkarshkrsingh/restapi/internal/routes"
)

func main() {
	router := mux.NewRouter()
	routes.HandleRecordRoutes(router)

	log.Println("server starting at port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("server failed: %v\n", err)
	}
}
