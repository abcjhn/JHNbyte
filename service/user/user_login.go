package usersvc



type UserLoginFlow struct{
	username string
	passtword string
}

func(u *UserLoginFlow)CheckParam() error{
	//用户名和密码规范逻辑
	return nil
}

func(u *UserLoginFlow)Login() (int,error){
	//判断用户是否存在
	//判断密码是否正确
	//登录状态
	return 0,nil
}
