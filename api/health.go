// api/health.go
package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Estrutura para resposta de saúde
type HealthResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// RegisterHealthRoutes registra todas as rotas relacionadas ao health check
func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthHandler)
}

// Função que trata requisições ao endpoint /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Define o status code como 200 OK
	statusCode := http.StatusOK

	// Registra nos logs quem consultou a API, o método e o status code
	log.Printf("Health check - IP: %s | Método: %s | Status Code: %d",
		r.RemoteAddr, r.Method, statusCode)

	// Define o tipo de conteúdo como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Cria a resposta incluindo o código de status
	response := HealthResponse{
		Status: "ok",
		Code:   statusCode,
	}

	// Codifica a resposta como JSON e escreve no ResponseWriter
	json.NewEncoder(w).Encode(response)
}
