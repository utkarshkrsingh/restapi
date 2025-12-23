package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRecordRoutes(router *mux.Router) {
	router.HandleFunc("/v1/movies", CreateRecord()).Methods(http.MethodPost)
	router.HandleFunc("/v1/movies", GetRecords()).Methods(http.MethodGet)
	router.HandleFunc("/v1/movies/{id}", UpdateRecords()).Methods(http.MethodPut)
	router.HandleFunc("/v1/movies/{id}", DeleteRecords()).Methods(http.MethodDelete)
}

func CreateRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create Record")
	}
}
func GetRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Record")

	}
}
func UpdateRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Update Record")

	}
}
func DeleteRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Delete Record")
	}
}
