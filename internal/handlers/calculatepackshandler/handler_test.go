package calculatepackshandler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"re_partners/internal/handlers/calculatepackshandler"
)

func TestHandler_GetHeaders(t *testing.T) {
	healthHandler := calculatepackshandler.New(nil)
	assert.Empty(t, healthHandler.GetHeaders())
}

func TestHandler_GetMethod(t *testing.T) {
	healthHandler := calculatepackshandler.New(nil)
	assert.Equal(t, http.MethodGet, healthHandler.GetMethod())
}

func TestHandler_GetPath(t *testing.T) {
	healthHandler := calculatepackshandler.New(nil)
	assert.Equal(t, "/packs", healthHandler.GetPath())
}
