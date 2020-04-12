package PersonController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fenriz07/api-go/person"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {

	person := person.NewPerson("Servio", "Zambrano", 28)

	output, err := json.Marshal(person)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Algo salio mal", err)
	}

	fmt.Fprint(w, string(output))
}
