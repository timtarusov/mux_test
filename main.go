package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	fmt.Println("Go to http://localhost:80")
	http.ListenAndServe(":80", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"first":1, "second":2})
}
