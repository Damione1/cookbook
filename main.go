package main

import (
	"fmt"
	"net/http"

	database "github.com/Damione1/portfolio-playground/db"
)

func main() {
	database.ConnectDb()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !!!")
	})
	http.ListenAndServe(":3000", nil)

}
