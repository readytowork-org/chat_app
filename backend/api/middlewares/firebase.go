package middlewares

import (
	"letschat/api/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FirebaseAuth struct {
	service services.FirebaseService
}

func NewFirebaseAuth(service services.FirebaseService) FirebaseAuth {
	return FirebaseAuth{
		service: service,
	}
}

func (m *FirebaseAuth) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationToken := ctx.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
			ctx.Abort()
			return
		}

		//verify token
		token, err := m.service.VerifyToken(idToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("UUID", token.UID)
		ctx.Next()
	}
}
