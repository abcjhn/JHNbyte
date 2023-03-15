package usersvc

import (
	"errors"
	"mydousheng/repository"
	"mydousheng/util"
)



type UserRegisterFlow struct{
	username string
	password string
}

func Register(name string,pwd string) (int,error){
	id,err := NewUserRegisterFlow(name,pwd).Do()
	if err!=nil{
		return 0,err
	}
	return id,nil

}

func NewUserRegisterFlow(name string,password string) *UserRegisterFlow{
	return &UserRegisterFlow{
		username: name,
		password: password,
	}
}

func (u *UserRegisterFlow)Do()(int,error)  {
	if err:=u.CheckParam();err!=nil{
		return 0,err
	} 
	id,err := u.Register();
	if err!=nil{
		return 0,err
	}
	return id,nil
	
}

func(u *UserRegisterFlow)CheckParam() error{
	
	return nil
}

func(u *UserRegisterFlow)Register() (int,error){
	//用户密码加密
	PasswordHash,err := util.PwdHash(u.password)
	if err!=nil{
		return 0,err
	}
	//用户是否存在
	userquery,err := repository.NewUserDaoInstance().QueryUserByName(u.username)
	if err!=nil{
		return 0,err
	}else if err==nil&&userquery.Id!= 0 {
		return 0,errors.New("用户已经存在")		
	}
	//创建用户
	userinfo := &repository.User{
		Username: u.username,
		Password: PasswordHash,
		IsDeleted: false,
	}
	if err:= repository.NewUserDaoInstance().CreateUser(userinfo);err!=nil{
		return 0,err
	}
	return 0,nil
}