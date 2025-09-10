package cyoa

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the Home Page XD")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Esta es la ágina del About u.u")
}

func StartingServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	log.Println("Server starting on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
