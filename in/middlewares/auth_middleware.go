package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getUserIDFromClaims(claims jwt.MapClaims) (int, bool) {
	uid, ok := claims["user_id"]
	if !ok {
		return 0, false
	}
	switch v := uid.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	default:
		return 0, false
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token não informado",
			})
			c.Abort()
			return
		}

		// Bearer TOKEN
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := []byte(os.Getenv("JWT_SECRET"))
		if len(secret) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET não configurado"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido ou expirado",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}
		userID, ok := getUserIDFromClaims(claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: user_id ausente"})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
