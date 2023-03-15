package usersvc

import (
	"errors"
	"mydousheng/repository"
	"mydousheng/util"
)



type UserLoginFlow struct{
	username string
	password string
}

func Login(username string, pwd string) (int,error) {
	return NewUserLoginFlow(username,pwd).Do()
	
}

func NewUserLoginFlow(username string ,pwd string) *UserLoginFlow {
	return &UserLoginFlow{
		username: username,
		password: pwd,
	}
}


func (u *UserLoginFlow)Do() (int,error)  {
	if err:= u.CheckParam();err!=nil{
		return 0,err
	}
	id,err := u.Login()
	if err !=nil{
		return 0,err
	}
	return id,nil
	
}

func(u *UserLoginFlow)CheckParam() error{
	//用户名和密码规范逻辑
	return nil
}

func(u *UserLoginFlow)Login() (int,error){
	//判断用户是否存在
	userquery,err := repository.NewUserDaoInstance().QueryUserByName(u.username)
	if err !=nil||userquery.Id ==0 {
		return 0,errors.New("用户不存在")
	}
	//判断密码是否正确
	if !util.PwdVerify(u.password,userquery.Password){
		return 0,errors.New("密码错误")
	}
	//登录状态
	return userquery.Id,nil
}
