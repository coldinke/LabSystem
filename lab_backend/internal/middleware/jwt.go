package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"lab_backend/internal/handler"
	"lab_backend/internal/reason"
	"net/http"
	"time"
)

var jwtSecret = []byte("hellltestjwt")

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expireTime.Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			handler.HandleResponse(ctx, errors.New(http.StatusUnauthorized, reason.UnauthorizedError), nil)
			ctx.Abort()
			return
		}

		tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			handler.HandleResponse(ctx, errors.New(http.StatusUnauthorized, reason.UnauthorizedError), nil)
			ctx.Abort()
			return
		}

		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			ctx.Set("username", claims["username"])
		} else {
			handler.HandleResponse(ctx, errors.New(http.StatusUnauthorized, reason.UnauthorizedError), nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
