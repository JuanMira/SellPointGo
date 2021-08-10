package auth

import (
	auth_query "gingonic/querys/auth"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
)

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaim struct{
	ID uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {

	var userBody *UserBody

	if err := c.ShouldBindBodyWith(&userBody, binding.JSON);err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"data invalid",			
		})		
		return		
	}	

	exist,err,user := auth_query.Login(userBody.Username,userBody.Password)	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})	
		return	
	}

	if exist == false{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"User isn't exist",
		})
		return
		
	}	

	claims := CustomClaim{
		ID: user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),						
		},
	}

	secrt := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secrt))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Cannot create token try it again",
		})
		return 
	}	

	c.JSON(http.StatusAccepted, gin.H{
		"messsage": "loged",
		"token":tokenString,
	})
	return
}
