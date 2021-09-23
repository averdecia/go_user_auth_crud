package middleware

import (
	"crud/config"
	"crud/entities"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTContent is used to save data on the JWT
type JWTContent struct {
	UserID primitive.ObjectID `json:"user_id"`
	jwt.StandardClaims
}

// CookieUserAuthorization checks for auth cookie
func CookieUserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization cookie"})
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &JWTContent{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT), nil
		})

		if sessionData, ok := token.Claims.(*JWTContent); ok && token.Valid {
			fmt.Printf("%v %v", sessionData.UserID, sessionData.StandardClaims.ExpiresAt)
			c.Set("Session", sessionData)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is expired"})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is invalid"})
			}
			fmt.Printf("jwt error %v", ve)

			return
		}

	}
}

// TokenUserAuthorization checks for auth basic
func TokenUserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.GetHeader("Authorization")
		elements := strings.Split(cookie, " ")

		token, err := jwt.ParseWithClaims(elements[1], &JWTContent{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT), nil
		})

		if sessionData, ok := token.Claims.(*JWTContent); ok && token.Valid {
			fmt.Printf("%v %v", sessionData.UserID, sessionData.StandardClaims.ExpiresAt)
			c.Set("Session", sessionData)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is expired"})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is invalid"})
			}
			fmt.Printf("jwt error %v", ve)

			return
		}

	}
}

// SetSessionToken static function in order to generate session cookie from user
func SetSessionToken(user *entities.User, c *gin.Context) (string, error) {
	content := JWTContent{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 15000,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, content)
	ss, err := token.SignedString([]byte(config.Config.JWT))
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}
	c.SetCookie("auth", ss, 15000, "/", "localhost", false, false)
	return ss, nil
}
