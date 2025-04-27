package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andresantiiago/go-api/api"
)

// Função que inicia a API web na porta 8080
func startAPI() error {
	// Cria um novo router
	mux := http.NewServeMux()

	// Registra as rotas de cada módulo
	api.RegisterHealthRoutes(mux)
	api.RegisterNamespaceRoutes(mux)

	// Inicia o servidor na porta 8080
	port := 8080
	log.Printf("Servidor iniciado na porta %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func main() {
	// Inicia a API e trata possíveis erros
	if err := startAPI(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
