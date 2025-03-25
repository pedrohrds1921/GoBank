package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("falied to encode json %w", err)
	}
	return nil

}
