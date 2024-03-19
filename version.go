package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Version number of auth application.
const Version = "1.0.0"

// version is the handler for the version endpoint.
func (a *Auth) version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"app": "auth", "version": Version})
}
