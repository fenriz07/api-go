package main

import (
	"fmt"
	"net/http"

	"github.com/fenriz07/api-go/controllers/PersonController"
)

func main() {

	http.HandleFunc("/api/person/1", PersonController.GetPerson)

	fmt.Println("Iniciando server")
	http.ListenAndServe(":8080", nil)
}
