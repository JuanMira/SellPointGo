package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyAdmin() gin.HandlerFunc {
	return func(c *gin.Context){
		secret := os.Getenv("SECRET_KEY")
		tokenH := c.GetHeader("Authorization")

		token, err := jwt.ParseWithClaims(
			tokenH,
			&CustomClaim{},
			func(t *jwt.Token) (interface{}, error) {
				return []byte(secret),nil
			},
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Token wrong pls login",				
			})								
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*CustomClaim)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Not authorized",
			})
			c.Abort()
			return
		}

		if claims.Role != true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"Don´t have the permissions",
			})
			c.Abort()
			return
		}

		c.Next()
		return
	}
}