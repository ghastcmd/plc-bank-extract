package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ghastcmd/plc-bank-extract/controllers"
)

func main() {
	http.HandleFunc("/", controllers.FetchMonthData)

	fmt.Printf("Starting server at %s\n", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil); err != nil {
		log.Fatal(err)
	}
}
