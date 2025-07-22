package routes_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/barely-sad-one/bonsai/internal/app/config"
	"github.com/barely-sad-one/bonsai/internal/app/server"
)

func TestPingRoute(t *testing.T) {
	config.InitConfig("../../../configs/server.json")
	app := server.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
