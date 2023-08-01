package errors

import (
	"net/http"
)

func InternalServer(w http.ResponseWriter, err *error) error {
	if *err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte((*err).Error()))
		return *err
	}
	return nil
}

func StatusBadRequest(w http.ResponseWriter, err *error) error {
	if *err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte((*err).Error()))
		return *err
	}
	return nil
}
