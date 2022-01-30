package httpServer

import (
	"apiServer/animal"
	"fmt"
	"net/http"
)

func Run() {
	fmt.Println("Starting server")
	animal.InitAnimalRoutes()
	http.ListenAndServe(":8080", nil)
}
