package handler

import (
	"net/http"

	"github.com/hoainam183/todo-app/pkg/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJson(w, 200, struct{}{})
}
