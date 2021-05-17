package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type customers struct {
	Name     string `json:"Name" xml:"Name"`
	City     string `json:"City" xml:"City"`
	Zip_code string `json:"Zip-Code" xml:"Zip-Code"`
}

func handle_f(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	Customers_l := []customers{{Name: "Priya", City: "jamshedpur", Zip_code: "831017"}, {Name: "Priya", City: "jamshedpur", Zip_code: "831017"}}

	fmt.Println(Customers_l)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(Customers_l)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(Customers_l)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
func getAllCustomers_JSON(w http.ResponseWriter, r *http.Request) {
	Customers_l := []customers{{Name: "Priya", City: "jamshedpur", Zip_code: "831017"}, {Name: "Priya", City: "jamshedpur", Zip_code: "831017"}}

	fmt.Println(Customers_l)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Customers_l)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getAllCustomers_XML(w http.ResponseWriter, r *http.Request) {
	Customers_l := []customers{{Name: "Priya", City: "jamshedpur", Zip_code: "831017"}, {Name: "Priya", City: "jamshedpur", Zip_code: "831017"}}

	fmt.Println(Customers_l)
	w.Header().Add("Content-Type", "application/xml")
	err := xml.NewEncoder(w).Encode(Customers_l)
	if err != nil {
		fmt.Println(err)
		return
	}
}
