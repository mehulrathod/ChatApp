package api

import (
	req "ChatApp/resources/request/api"
	res "ChatApp/resources/response/api"
	"database/sql"
	"github.com/sirupsen/logrus"
)

const (
	findOneUser            = "SELECT id, email, mobile, created_at, first_name, last_name FROM user WHERE id = ? AND deleted_at IS NULL"
	insertUser             = "INSERT INTO user(id, email, mobile, password, first_name, last_name) VALUES(?,?,?,?,?,?)"
	getUserCount           = "SELECT COUNT(*) FROM user WHERE email = ? AND deleted_at IS NULL"
)

type UserRepository interface {
	UserSignUp(request *req.SignUpRequest) (*res.UserResponse, error)
	CheckUser(email string) (int, error)
}

type userRepo struct {
	DB *sql.DB
}

func NewUserWriter(db *sql.DB) UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (ur *userRepo) UserSignUp(user *req.SignUpRequest) (*res.UserResponse, error) {
	logrus.Info("INFO : ", "User Repo Called(SignUp).")
	tx, err := ur.DB.Begin()
	if err != nil {
		logrus.Error("ERROR(tx) : ", err.Error())
		return &res.UserResponse{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Error("ERROR(tx) : ", err.Error())
		}
	}()

	_, err = tx.Exec(insertUser,user.ID, user.Email, user.Mobile, user.Password, user.FirstName, user.LastName)
	if err != nil {
		logrus.Error("ERROR(query) : ", err.Error())
		return &res.UserResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Error("ERROR(tx) : ", err.Error())
		return &res.UserResponse{}, err
	}

	userData, err := ur.getOneUser()
	if err != nil {
		logrus.Error("ERROR(data) : ", err.Error())
		return &res.UserResponse{}, err
	}

	return userData, nil
}

func (ur *userRepo) getOneUser() (*res.UserResponse, error) {
	user := &res.UserResponse{}
	err := ur.DB.QueryRow(findOneUser).Scan(&user.ID, &user.Email, &user.Mobile, &user.CreatedAt, &user.FirstName, &user.LastName)
	if err != nil {
		logrus.Error("ERROR(query) : ", err.Error())
		return &res.UserResponse{}, err
	}
	return user, nil
}

func (ur *userRepo) CheckUser(email string) (int, error) {
	var userCount int
	row := ur.DB.QueryRow(getUserCount, email)
	err := row.Scan(&userCount)
	if err != nil {
		logrus.Error("ERROR(query) : ", err.Error())
		return 0, err
	}
	return userCount, nil
}