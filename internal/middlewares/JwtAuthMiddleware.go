package middlewares

import (
	"net/http"

	"github.com/Bonk123123/notes_server/internal/config"
	"github.com/Bonk123123/notes_server/internal/pkg"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        err := pkg.ValidateJWT(c, config.ACCESS)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            c.Abort()
            return
        }
        c.Next()
    }
}