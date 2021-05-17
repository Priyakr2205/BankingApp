package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	hmux := mux.NewRouter() // reating new multpipler router
	hmux.HandleFunc("/greet", handle_f)
	hmux.HandleFunc("/customers_json", getAllCustomers_JSON)
	hmux.HandleFunc("/customers", getAllCustomers)
	hmux.HandleFunc("/customers_xml", getAllCustomers_XML)

	hmux.HandleFunc("/customers/{customer_id}", getCustomerId)

	http.ListenAndServe("localhost:8000", hmux) // passing new router , To use default use nil in 2nd param
}
