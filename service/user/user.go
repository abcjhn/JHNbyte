package usersvc

import "mydousheng/repository"

type UserInfo struct{
	Id int 			`json:"id"`
	name string		`json:"name"`
}

func QueryUserInfo(id int) (*UserInfo,error) {
	
	user,err := repository.NewUserDaoInstance().QueryUserById(id)
	if err!=nil{
		return nil,err
	}
	return &UserInfo{
		Id: user.Id,
		name: user.Username,
	},nil
}