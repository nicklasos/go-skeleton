package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type response map[string]interface{}

func bodyMarshal(w http.ResponseWriter, x map[string]interface{}) error {
	resp, err := json.Marshal(x)
	if err != nil {
		w.Write([]byte(`{"error": "Internal error"}`))
		return errors.New("An internal error has occurred.")
	}

	w.Write(resp)
	return nil
}
