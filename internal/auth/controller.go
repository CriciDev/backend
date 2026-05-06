package auth

import (
	"encoding/json"
	"net/http"

	"github.com/CriciumaDevJobs/backend/handlers"
)

type AuthenticationController struct {
	AuthUseCase *AuthenticationUseCase
}

func NewAuthenticationController(usecase *AuthenticationUseCase) *AuthenticationController {
	auth := AuthenticationController{
		AuthUseCase: usecase,
	}

	return &auth
}

func (controller *AuthenticationController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var req = AuthenticationRequest{}

	jsonErr := json.NewDecoder(r.Body).Decode(&req)

	if jsonErr != nil {
		handlers.ResponseWithHttpError(w, http.StatusBadRequest, "JSON Enviado não segue o padrão esperado")
		return
	}

	resp, httpErr := controller.AuthUseCase.AuthenticateUser(r.Context(), req.Email, req.Password)

	if httpErr != nil {
		handlers.ResponseWithHttpError(w, httpErr.Code, httpErr.Message)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
