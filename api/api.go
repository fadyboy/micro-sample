package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type API struct {
	Logger *log.Logger
}

// NewAPI is a constructor for the API
func NewAPI(log *log.Logger) *API {
	return &API{Logger: log}
}

func (a *API) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, your request is: %s", r.URL.Path)
}

func (a *API) Health(w http.ResponseWriter, r *http.Request) {
	a.Logger.Println("Checking service is up")

	status := map[string]string{
		"status": "UP",
	}

	json.NewEncoder(w).Encode(status)
}
