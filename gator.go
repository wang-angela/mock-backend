package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Gator struct {
	Name     string `json:"name"`
	Standing string `json:"standing"`
	GradYear int    `json:"grad_year"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/albert", exampleHandler).Methods("GET")

	err := http.ListenAndServe(":3031", r)
	if err != nil {
		log.Fatal(err)
	}
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	albert := Gator{
		Name:     "Albert",
		Standing: "Senior",
		GradYear: 2024}

	err := json.NewEncoder(w).Encode(&albert)
	if err != nil {
		log.Fatal(err)
	}

}
