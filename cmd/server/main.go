package main

import (
	"fmt"
	"net/http"

	"github.com/CriciumaDevJobs/backend/infra"
	"github.com/CriciumaDevJobs/backend/internal/devs"
)

func main() {

	db := infra.InitDB()

	devs.InitializeDevContext(db)

	fmt.Println("Iniciando servidor")

	fmt.Println("Servidor iniciado")
	http.ListenAndServe(":8080", nil)
}
