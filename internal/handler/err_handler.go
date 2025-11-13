package handler

import (
	"net/http"

	"github.com/hoainam183/todo-app/pkg/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithError(w, 400, "some thing went wrong")
}
