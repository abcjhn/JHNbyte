package repository

import (
	"sync"
	"time"
)

type User struct{
	Id			int			`gorm:"column:id" gorm:"primaryKey"`
	Username 	string 		`gorm:"column:username" gorm:"uniqueIndex"`
	Password 	string		`gorm:"column:password"`
	CreatedAt	time.Time	`gorm:"column:created_time"`
	UpdatedAt 	time.Time	`gorm:"column:updated_time"`
	IsDeleted 	bool		`gorm:"column:is_deleted"`
}
func (u User) TableName() string {
	return "user"
}

type UserDao struct{}
var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao{
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}


func ( *UserDao)CreateUser(user *User)  error{
	if err := GetDb().Create(user).Error; err!=nil{
		return err
	}
	return nil
}

func ( *UserDao)QueryUserById(id int) (*User,error){
	var user User
	if err := GetDb().Where("id=?",id).Find(&user).Error; err!=nil{
		return nil,err
	}
	return &user,nil
}

func ( *UserDao)QueryUserByName(name string)  (*User,error){
	var user User
	if err := GetDb().Where("username=?",name).Find(&user).Error;err!=nil{
		return nil,err
	}
	return &user,nil
}

