package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/endpoint", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	for i := 1; i <= 200; i++ {
		response := Response{
			Message: "Response " + strconv.Itoa(i),
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed to marshal JSON response:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)
		w.Write([]byte("\n"))
	}
}
