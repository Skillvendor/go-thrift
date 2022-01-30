package animal

import "net/http"

func InitAnimalRoutes() {
	http.HandleFunc("/api/animals", MyFunc)
}
