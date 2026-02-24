package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sviatilnik/sso/internal/sso/application"
	"github.com/sviatilnik/sso/internal/sso/application/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loginRequest := new(application.LoginRequest)

	err = json.Unmarshal(rawBody, loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginResponse, err := h.authService.Login(r.Context(), loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to serialize login response to json"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
