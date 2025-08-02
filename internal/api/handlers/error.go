package handlers

import (
	"net/http"
	"welp/internal/api/model/werror"
)

func Error(w http.ResponseWriter, err werror.Error) {
	http.Error(w, err.Msg, err.Code)
}
