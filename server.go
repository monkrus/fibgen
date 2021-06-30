package server

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server up")
}
