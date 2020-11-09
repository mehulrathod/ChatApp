package models

type User struct {
	Model
	FirstName string `gorm:"varchar(40)" json:"first_name" validate:"required"`
	LastName  string `gorm:"varchar(40)" json:"last_name" validate:"required"`
	Email     string `gorm:"varchar(50)" json:"email" validate:"required,email"`
	Password  string `gorm:"varchar(255)" json:"password" validate:"required"`
	Mobile    string `gorm:"varchar(15)" json:"mobile" validate:"required"`
}

func (u *User) TableName() string {
	return "user"
}
