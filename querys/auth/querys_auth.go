package auth

import (
	"fmt"
	"gingonic/core"
	"gingonic/services"

	"time"
)

func Login(username string, password string) (bool,error,core.User) {
	db,err:=services.ConnectDB()

	var user core.User

	if err != nil {
		return false, err, user
	}	

	userN := db.Where("username = ?",username).First(&user)

	same := services.ComparePassword(password,user.Password)

	if same == false {
		return false, nil , user
	}

	if userN == nil {
		return false, nil , user
	}	
	
	user.Password = ""

	return true, nil,user
}

func Register(user core.UserBody_R)(bool,error){
	
	var user_db core.User = core.User {
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		Birthday: user.Birthday,
		LastLogin: time.Now(),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Phone: user.Phone,
		CountryId: user.CountryId,
		BillingAddress: user.BillingAddress,						
	}	

	db, err := services.ConnectDB()	

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connecting db")
		return false, err
	}

	result := db.Create(&user_db)		

	fmt.Print("Rows affected = ")
	fmt.Println(result.RowsAffected)

	return true, nil
}