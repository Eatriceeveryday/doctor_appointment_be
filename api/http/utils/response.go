package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
