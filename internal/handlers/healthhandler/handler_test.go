package healthhandler_test

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"

	"re_partners/internal/handlers/healthhandler"
)

func TestHandler_GetHeaders(t *testing.T) {
	healthHandler := healthhandler.New()
	assert.Empty(t, healthHandler.GetHeaders())
}

func TestHandler_GetMethod(t *testing.T) {
	healthHandler := healthhandler.New()
	assert.Equal(t, http.MethodGet, healthHandler.GetMethod())
}

func TestHandler_GetPath(t *testing.T) {
	healthHandler := healthhandler.New()
	assert.Equal(t, "/health", healthHandler.GetPath())
}

func TestHandler_ServeHTTP(t *testing.T) {
	healthHandler := healthhandler.New()

	apitest.New().
		Handler(healthHandler).
		Get(healthHandler.GetPath()).
		Expect(t).
		Status(http.StatusOK).
		Body(`{"status":"OK"}`).
		End()
}
