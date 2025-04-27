// api/namespace.go
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Namespace representa a estrutura de dados do namespace
type Namespace struct {
	Name       string            `json:"name"`
	Labels     map[string]string `json:"labels"`
	AdminUsers []string          `json:"admin_user"`
	AdminGroup []string          `json:"admin_group"`
	Flavor     string            `json:"flavor"`
}

// RegisterNamespaceRoutes registra todas as rotas relacionadas a namespaces
func RegisterNamespaceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/namespace", namespaceHandler)
}

// Função que trata requisições ao endpoint /namespace
func namespaceHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		statusCode := http.StatusMethodNotAllowed

		log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Mensagem: Método não permitido",
			r.RemoteAddr, r.Method, statusCode)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Apenas o método POST é permitido para este endpoint",
		})
		return
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		statusCode := http.StatusBadRequest

		log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Erro: %v",
			r.RemoteAddr, r.Method, statusCode, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Erro ao ler o corpo da requisição",
		})
		return
	}

	// Decodifica o JSON para a estrutura Namespace
	var namespace Namespace
	if err := json.Unmarshal(body, &namespace); err != nil {
		statusCode := http.StatusBadRequest

		log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Erro: %v",
			r.RemoteAddr, r.Method, statusCode, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JSON inválido ou incompatível com o modelo esperado",
		})
		return
	}

	// Validações de todos os campos
	validationError := validateNamespace(namespace)
	if validationError != "" {
		statusCode := http.StatusBadRequest

		log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Erro: %s",
			r.RemoteAddr, r.Method, statusCode, validationError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": validationError,
		})
		return
	}

	// Chamando função para fazer o commit do namespace
	commitError := commitNamespace(namespace)
	if commitError != nil {
		statusCode := http.StatusInternalServerError

		log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Erro: %v",
			r.RemoteAddr, r.Method, statusCode, commitError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Erro ao realizar o commit do namespace",
		})
		return
	}

	// Resposta de sucesso
	statusCode := http.StatusCreated

	log.Printf("Namespace - IP: %s | Método: %s | Status Code: %d | Namespace: %s",
		r.RemoteAddr, r.Method, statusCode, namespace.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "success",
		"code":      statusCode,
		"message":   "Namespace criado com sucesso",
		"namespace": namespace,
	})
}

// validateNamespace verifica se todos os campos do namespace são válidos
func validateNamespace(ns Namespace) string {
	// Valida o nome
	if ns.Name == "" {
		return "O nome do namespace não pode estar vazio"
	}

	// Valida as labels
	if len(ns.Labels) == 0 {
		return "O campo labels não pode estar vazio"
	}

	// Verifica se alguma label está vazia
	for key, value := range ns.Labels {
		if key == "" {
			return "As chaves das labels não podem estar vazias"
		}
		if value == "" {
			return fmt.Sprintf("O valor da label '%s' não pode estar vazio", key)
		}
	}

	// Valida AdminUsers
	if len(ns.AdminUsers) == 0 {
		return "O campo admin_user não pode estar vazio"
	}

	// Verifica se algum usuário admin está vazio
	for i, user := range ns.AdminUsers {
		if user == "" {
			return fmt.Sprintf("O usuário admin na posição %d não pode estar vazio", i)
		}
	}

	// Valida AdminGroup
	if len(ns.AdminGroup) == 0 {
		return "O campo admin_group não pode estar vazio"
	}

	// Verifica se algum grupo admin está vazio
	for i, group := range ns.AdminGroup {
		if group == "" {
			return fmt.Sprintf("O grupo admin na posição %d não pode estar vazio", i)
		}
	}

	// Valida o Flavor
	if ns.Flavor == "" {
		return "O campo flavor não pode estar vazio"
	}

	return "" // Sem erros
}

func commitNamespace(ns Namespace) error {
	return nil
}
