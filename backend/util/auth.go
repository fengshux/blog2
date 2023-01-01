package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fengshux/blog2/backend/conf"
	"github.com/fengshux/blog2/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user *model.User) (string, error) {

	c := conf.GetConf().Auth
	var sampleSecretKey = []byte(c.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Duration(c.Expires) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func authMiddleware(hard, admin bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		user, err := extractClaims(tokenString)

		if user != nil && err == nil {
			// 如果是要求admin权限，
			if admin && user.Role != model.USER_ROLE_ADMIN {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "没有权限"})
				return
			}
			c.Set("userId", user.ID)
			c.Set("role", user.Role)

		} else if hard {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "请登录"})
			return
		}

		c.Next()
	}
}

func SoftAuth() gin.HandlerFunc {
	return authMiddleware(false, false)
}

func HardAuth() gin.HandlerFunc {
	return authMiddleware(true, false)
}

func AdminAuth() gin.HandlerFunc {
	return authMiddleware(true, true)
}

func extractClaims(tokenString string) (*model.User, error) {
	c := conf.GetConf().Auth
	var sampleSecretKey = []byte(c.Secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}

		return sampleSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return &model.User{
			ID:   int64(userId),
			Role: role,
		}, nil
	}
	return nil, fmt.Errorf("token is not valid or expired")
}
