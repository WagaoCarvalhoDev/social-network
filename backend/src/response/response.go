package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func ToJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Err(w http.ResponseWriter, statusCode int, err error) {
	ToJson(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}
