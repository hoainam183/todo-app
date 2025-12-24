package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hoainam183/todo-app/internal/services"
	"github.com/hoainam183/todo-app/pkg/common/apperrors"
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

		err := svc.Register(req.Username, req.Password)
		if err != nil {
			utils.ResponseWithError(w, getStatusCode(err), err.Error())
			return
		}

		utils.ResponseWithJson(w, http.StatusCreated, map[string]string{
			"message": "User registered successfully",
		})
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

// LoginHandler returns an http handler that performs user login and returns JWT token.
func LoginHandler(svc *services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		res, err := svc.Login(req.Username, req.Password)
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		utils.ResponseWithJson(w, http.StatusOK, loginResponse{
			UserId: res.UserId,
			Token:  res.Token,
		})
	}
}

func getStatusCode(err error) int {
	switch {
	case errors.Is(err, apperrors.ErrUserExists):
		return http.StatusConflict
	case errors.Is(err, apperrors.ErrMissingField):
		return http.StatusBadRequest
	default:
		return http.StatusBadRequest
	}
}
