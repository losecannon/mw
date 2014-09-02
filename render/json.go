package render

import (
	"encoding/json"
	"net/http"
)

func JSON(rw http.ResponseWriter, status int, data interface{}) {
	var result []byte
	var err error

	result, err = json.Marshal(data)
	if err != nil {
		http.Error(rw, err.Error(), 500)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(result)
}
