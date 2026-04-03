package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	msg := Response{Message: "Message from net/http"}

	jsonData, err := json.Marshal(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)

}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}
