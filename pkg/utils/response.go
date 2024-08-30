package utils

import (
	"encoding/json"
	"net/http"

	"github.com/mhshajib/oasis_boilerplate/pkg/config"
)

// Response ...
type Response struct {
	Status     int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Error      interface{} `json:"error,omitempty"` // this field will be ommited from user response body based on log level
}

// Render ...
func (r *Response) Render(w http.ResponseWriter) error {
	if !config.HttpApp().Verbose {
		r.Error = nil // if verbose set to false then remove the error from public response
	}
	bb, err := json.Marshal(r)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}
	_, err = w.Write(bb)
	return err
}

// M represents a generic map
type M map[string]interface{}
