package usersvc


type UserRegisterFlow struct{
	username string
	password string
}

func(u *UserRegisterFlow) NewUserRegister(name string,password string) *UserRegisterFlow{
	return &UserRegisterFlow{
		username: name,
		password: password,
	}
}

func(u *UserRegisterFlow)CheckParam() error{
	//补充用户名和密码规范
	return nil
}

func(u *UserRegisterFlow)Register() (int,error){
	//用户密码校对
	//用户是否存在
	//创建用户
	return 0,nil
}