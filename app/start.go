package app

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", handle_f)
	http.HandleFunc("/customers_json", getAllCustomers_JSON)
	http.HandleFunc("/customers", getAllCustomers)
	http.HandleFunc("/customers_xml", getAllCustomers_XML)

	http.ListenAndServe("localhost:8000", nil)
}
