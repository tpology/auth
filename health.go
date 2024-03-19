package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthz is the handler for the healthz endpoint.
func (a *Auth) healthz(c *gin.Context) {
	c.Status(http.StatusOK)
}

// readyz is the handler for the readyz endpoint.
func (a *Auth) readyz(c *gin.Context) {
	c.Status(http.StatusOK)
}
