package utils

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(UserID uint) (string, error) {

	token_lifespan := 3600

	claims := jwt.MapClaims{}
	claims["UserID"] = UserID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("HEHETOKEN"))

}

func TokenValid(c *gin.Context) int {
	tokenString := ExtractToken(c)
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("MUTHUS-TOKEN"), nil
	})
	if err != nil {
		return 0
	}
	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		userId := int(claims["UserID"].(float64))
		return userId
	}
	if err != nil {
		return 0
	}
	return 0
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
