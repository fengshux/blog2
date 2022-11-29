package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fengshux/blog2/backend/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// TODO Secret inject in environment variable

func GenerateJWT(userId int64) (string, error) {

	c := conf.GetConf().Auth
	var sampleSecretKey = []byte(c.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(c.Expires) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func authMiddleware(hard bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authentication")
		userId, err := extractClaims(tokenString)
		if userId != 0 && err == nil {
			c.Set("userId", userId)
		} else if hard {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "请登录"})
			return
		}

		c.Next()
	}
}

func SoftAuth() gin.HandlerFunc {
	return authMiddleware(false)
}

func HardAuth() gin.HandlerFunc {
	return authMiddleware(true)
}

func extractClaims(tokenString string) (int64, error) {
	c := conf.GetConf().Auth
	var sampleSecretKey = []byte(c.Secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}

		return sampleSecretKey, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := claims["userId"].(float64)
		return int64(userId), nil
	}
	return 0, fmt.Errorf("token is not valid or expired")
}
