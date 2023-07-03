package handler

import (
	"encoding/json"
	"net/http"
)

type RespFormat struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, code int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	successHttpStatus := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent,
	}

	isOk := false
	for _, value := range successHttpStatus {
		if value == code {
			isOk = true
			break
		}
	}

	return json.NewEncoder(w).Encode(RespFormat{
		Ok:      isOk,
		Message: message,
		Data:    data,
	})
}
