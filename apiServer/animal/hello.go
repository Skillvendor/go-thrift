package animal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pedigree struct {
	Identification string
}

type Human struct {
	Name string
	Age  int
}

type Animal struct {
	Name  string
	Age   int
	Owner []Human
	Misto Pedigree
}

func MyFunc(w http.ResponseWriter, r *http.Request) {
	myAnimal := Animal{}
	err := json.NewDecoder(r.Body).Decode(&myAnimal)

	if err != nil {
		fmt.Println("ERROR", err)
	}

	err = json.NewEncoder(w).Encode(myAnimal)

	if err != nil {
		fmt.Println("ERROR2", err)
	}
	// w.Write([]byte(myAnimal.Owner[0].Name))
}
