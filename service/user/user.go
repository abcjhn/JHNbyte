package usersvc

import "mydousheng/repository"

type UserInfo struct{
	Id int
	name string
}

func (u *UserInfo)QueryUserInfo(id int) (*UserInfo,error) {
	
	user,err := repository.Userdao.QueryUserById(id)
	if err!=nil{
		return nil,err
	}
	return &UserInfo{
		Id: user.Id,
		name: user.Username,
	},nil
}