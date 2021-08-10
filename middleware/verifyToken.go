package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaim struct{
	ID uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyToken() gin.HandlerFunc {
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

			//debug
			fmt.Println(err)			
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*CustomClaim)

		if !ok{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Not authorized",
			})
			c.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().UTC().Unix(){
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"token expires",
			})
			c.Abort()
			return
		}

		c.Next()
		return		

	}
}