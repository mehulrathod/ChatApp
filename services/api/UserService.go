package api

import (
	"ChatApp/models"
	repo "ChatApp/repository/api"
	req "ChatApp/resources/request/api"
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	SignUpUser(req *req.SignUpRequest) map[string]interface{}
}

type userService struct {
	User     models.User
	UserRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (us *userService) SignUpUser(req *req.SignUpRequest) map[string]interface{}  {
	userCount, err := us.UserRepo.CheckUser(req.Email)
	if err != nil {
		logrus.Error("ERROR(DB) : ", err.Error())
	}
	if userCount > 0 {
		logrus.Warn("WARN : ", "Duplicate User... for email ", req.Email)
	}
	resp, err := us.UserRepo.UserSignUp(req)
	 if err != nil {
	 	logrus.Error("ERROR(from repo) : ", err.Error())
	 }

	 fmt.Printf("SignUpSuccess %s" , resp)
	 return nil
}