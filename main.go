package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fenriz07/api-go/person"
)

func main() {

	http.HandleFunc("/api/person", func(w http.ResponseWriter, r *http.Request) {

		person := person.NewPerson("Servio", "Zambrano", 28)

		output, err := json.Marshal(person)

		if err != nil {
			fmt.Println("Algo salio mal", err)
		}

		fmt.Fprint(w, string(output))
	})

	fmt.Println("Iniciando server")
	http.ListenAndServe(":8080", nil)
}
