package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	var maxBytes int64 = 1_024_000
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(dst)
	if err != nil {
		fmt.Println(err)
		return errors.New("invalid request body")
	}

	return nil
}
