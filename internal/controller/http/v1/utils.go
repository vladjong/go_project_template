package v1

import (
	"encoding/json"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Error struct {
	Error string `json:"Error"`
}

func WriteError(w http.ResponseWriter, error error, status int) {
	w.WriteHeader(status)

	e := &Error{
		Error: error.Error(),
	}

	byteMessage, _ := json.Marshal(e)

	if _, err := w.Write(byteMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteResponseJson(w http.ResponseWriter, body interface{}) {
	resp, err := jsoniter.Marshal(body)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}
}
