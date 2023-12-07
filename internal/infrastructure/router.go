package infrastructure

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler ...
type Handler interface {
	GetPath() string
	GetMethod() string
	GetHeaders() map[string]string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// LoadRouter ...
func LoadRouter(handlers ...Handler) *mux.Router {
	router := mux.NewRouter()

	for _, handler := range handlers {
		router.
			Handle(handler.GetPath(), handler).
			Methods(handler.GetMethod()).
			Headers(func() []string {
				headersMap := handler.GetHeaders()
				var headers []string
				for key, value := range headersMap {
					headers = append(headers, key)
					headers = append(headers, value)
				}
				return headers
			}()...)
	}

	return router
}
