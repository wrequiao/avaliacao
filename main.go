package main

import (
	"avaliacao/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando app...")
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
