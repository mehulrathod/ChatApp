package routers

import (
	con "ChatApp/controllers/api"
	"ChatApp/models"
	repo "ChatApp/repository/api"
	serv "ChatApp/services/api"
)

func hUser() *con.UserController {
	rpo := repo.NewUserWriter(models.GetDB())
	as := serv.NewUserService(rpo)
	ac := con.UserController{
		UserService: as,
	}
	return &ac
}