package api

import (
	req "ChatApp/resources/request/api"
	s "ChatApp/services/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	UserService s.UserService
}

func (uc *UserController) SignUp(c *gin.Context) {
	logrus.Info("INFO : ", "User Controller Called(SignUp).")
	var user req.SignUpRequest

	//decode the request body into struct and failed if any error occur
	if err := c.BindJSON(&user); err != nil {
		logrus.Error("ERROR : ", err.Error())
		return
	}

	//call service
	resp := uc.UserService.SignUpUser(&user)

	fmt.Printf("Status Code %s",resp)
	logrus.Info("INFO : ", "Sign up complete...", resp["data"])
}
