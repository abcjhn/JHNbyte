package repository

import (
	"os/user"
	"time"
)

type User struct{
	Id			int			`gorm:"user_id"`
	Username 	string 		`gorm:"username"`
	Password 	string		`gorm:"password"`
	Create_time 	time.Time	`gorm:"create_time"`
	Updatetime 	time.Time	`gorm:"update_time"`
	IsDeleted 	bool		`gorm:"is_deleted"`
}
func (u User) TableName() string {
	return "user"
}

type UserDao struct{}
var userdao *UserDao

func (userdao *UserDao)CreateUser(user *User)  error{
	if err := db.Create(&user).Error; err!=nil{
		return err
	}
	return nil
}

func (userdao *UserDao)QueryUserById(i int) (*User,error){
	if err :=

}