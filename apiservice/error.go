package apiservice

import (
	"encoding/json"
	"log"
	"net/http"
)

type structuredError struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func newStructuredError(t string, err error) structuredError {
	return structuredError{
		Type:        t,
		Description: err.Error(),
	}
}

func sendError(w http.ResponseWriter, t string, err error, s int) {
	sErr := newStructuredError(t, err)
	js, _ := json.Marshal(sErr)

	w.WriteHeader(s)
	w.Write(js)
}

func ReturnBadRequestError(w http.ResponseWriter, t string, err error) {
	sendError(w, t, err, http.StatusBadRequest)
}

func ReturnNotFoundError(w http.ResponseWriter, t string, err error) {
	sendError(w, t, err, http.StatusNotFound)
}

func ReturnInternalServerError(w http.ResponseWriter, t string, err error) {
	log.Println(err)
	sendError(w, t, err, http.StatusInternalServerError)
}
