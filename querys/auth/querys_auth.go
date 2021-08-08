package auth

import "gingonic/services"

func Login(string username, string password) (bool,error) {
	db,err:=services.ConnectDB()
	
	if err != nil {
		return false, err
	}

	

}