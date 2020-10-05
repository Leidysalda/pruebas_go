package main

import (
	"fmt"
	"time"

	"meeting/models"

	"github.com/google/uuid"
)

func main() {
	users := models.User{
		Username:   "Lola",
		UserID:     uuid.New().String(),
		Email:      " ",
		CompanyID:  " ",
		CreatedAt:  time.Now().Format("2006-01-02T15:04:05-0700"),
		ModifiedAt: time.Now().Format("2006-01-02T15:04:05-0700"),
	}

	fmt.Println(users)

	//userJson := users.ToStruct()
	//fmt.Println(userJson)

	//fmt.Println(object)
	fmt.Println("hello world!")

	var (
		username string = "Bobby"
		email    string = "b@gmail.com"
		/**companyID string = "7489eue"
		isActive       bool   = true
		userID         string = "q5517h"
		invitationSend int    = 1
		passw          string = "****"
		create         string = time.Now().Format("2006-01-02T15:04:05-0700")
		modificate     string = time.Now().Format("2006-01-02T15:04:05-0700")*/
	)

	var user1, _ = (models.CreateUser(username, email))
	fmt.Println(user1)
}
