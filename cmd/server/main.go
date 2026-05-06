package main

import (
	"log"
	"net/http"

	"github.com/CriciumaDevJobs/backend/infra"
	"github.com/CriciumaDevJobs/backend/internal"
)

var (
	PORT = ":8080"
)

func main() {

	db := infra.InitDB()
	log.Printf("Conexão com Banco de Dados bem sucedida!")

	app := internal.StartAppContext(db)

	log.Printf("Subindo Servidor HTTP na Porta %s", PORT)
	err := http.ListenAndServe(PORT, app.Router)

	if err != nil {
		log.Printf("ERRO: Falha ao subir servidor HTTP! Message: %s", err.Error())
		panic(err)
	}
}
