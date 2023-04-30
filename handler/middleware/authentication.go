package middleware

import (
	"fmt"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func parseToken(tokenString string) (float64, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return 0, errs.NewUnauthorized("Token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		token, ok := claims["userId"].(float64)
		if !ok {
			return 0, errs.NewBadRequest("Failed to get userId from token claims")
		}
		return token, nil
	}

	return 0, errs.NewUnauthorized("Token is not valid")
}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			newError := errs.NewUnauthorized("Token should not be empty")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		tokenString := strings.Fields(header)

		if tokenString[0] != "Bearer" {
			newError := errs.NewUnauthorized("Token type must be a bearer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		userId, err := parseToken(tokenString[1])
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userId", uint(userId))

		ctx.Next()
	}
}
