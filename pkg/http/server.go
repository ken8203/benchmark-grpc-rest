package httprest

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func RunServer() {
	srv := &http.Server{
		Addr: ":8001",
	}

	http.HandleFunc("/", GreeterHandler)
	log.Fatalln(srv.ListenAndServe())
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	response := &Response{
		Message: fmt.Sprintf("Hey there %s %f", name, rand.Float64()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
