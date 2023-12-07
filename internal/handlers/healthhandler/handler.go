package healthhandler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (handler *Handler) GetHeaders() map[string]string {
	return map[string]string{}
}

func (handler *Handler) GetMethod() string {
	return http.MethodGet
}

func (handler *Handler) GetPath() string {
	return "/health"
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	body := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	bodyJSON, _ := json.Marshal(body)

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(bodyJSON)
}
