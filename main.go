package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func fetchMonthData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not suported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Bem vindo a rota de teste")
}

func main() {
	http.HandleFunc("/", fetchMonthData)

	fmt.Printf("Starting server at %s\n", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil); err != nil {
		log.Fatal(err)
	}
}
