package auth_query

import (
	"gingonic/core"
	"gingonic/services"
)

func Login(username string, password string) (bool,error,core.User) {
	db,err:=services.ConnectDB()

	var user core.User

	if err != nil {
		return false, err, user
	}	

	userN := db.Where("username = ?",username).First(&user)

	if userN == nil {
		return false, nil , user
	}

	passN := db.Where("password = ?",password).First(&user)

	if passN == nil {
		return false, nil , user
	}
			
	return true, nil,user
}