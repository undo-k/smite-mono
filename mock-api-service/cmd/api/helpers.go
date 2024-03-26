package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
// 	maxBytes := 1048576 // One megabyte
// 	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

// 	dec := json.NewDecoder(r.Body)
// 	err := dec.Decode(data)
// 	if err != nil {
// 		return err
// 	}

// 	err = dec.Decode(&struct{}{})
// 	if err != io.EOF {
// 		return errors.New("body must have only a single JSON value")
// 	}

// 	return nil
// }

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	output, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		return err
	}

	return nil
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	minutes := d.Duration / time.Minute
	seconds := (d.Duration % time.Minute) / time.Second
	return json.Marshal(fmt.Sprintf("%02d:%02d", minutes, seconds))
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

// func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
// 	statusCode := http.StatusBadRequest

// 	if len(status) > 0 {
// 		statusCode = status[0]
// 	}

// 	var payload jsonResponse
// 	payload.Error = true
// 	payload.Message = err.Error()

// 	return app.writeJSON(w, statusCode, payload)
// }
