package rest

import (
	"encoding/json"
	"net/http"

	"github.com/hoainam183/todo-app/internal/services"
	"github.com/hoainam183/todo-app/pkg/utils"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterHandler returns an http handler that performs user registration.
func RegisterHandler(svc *services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		user, err := svc.Register(req.Username, req.Password)
		if err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.ResponseWithJson(w, http.StatusCreated, map[string]string{
			"id":       user.ID,
			"username": user.Username,
		})
	}
}
