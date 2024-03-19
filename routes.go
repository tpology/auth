package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routes returns the routes for the auth application.
func (a *Auth) routes() http.Handler {
	g := gin.New()
	g.Use(gin.Recovery())
	g.GET("/healthz", a.healthz)
	g.GET("/readyz", a.readyz)
	g.GET("/version", a.version)
	return g
}
