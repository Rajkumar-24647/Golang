package main

import (
	"fmt"
	"github.com/Rajkumar-24647/Golang/auth"
	"github.com/Rajkumar-24647/Golang/user"
)

func main() {
	auth.LoginWithCredentials("rajkumar", "wwwvferr")

session := auth.GetSession()
fmt.Println(session)


user := user.User{
	Email: "raj@xyz.com",
	Name: "rajkumar",
}

fmt.Println(user)


}