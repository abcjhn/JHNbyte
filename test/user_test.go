package test

import (
	"fmt"
	"mydousheng/repository"
	"testing"
)
func TestCreateUser(t *testing.T) {

	if err:=repository.Init();err!=nil{
		return 
	}
	db :=repository.Db
	db.Migrator().CreateTable(&repository.User{})
	userdao := repository.Userdao
	userdao.CreateUser(&repository.User{Username:"jhnjhn"})


}
func TestQueryUserById(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return 
	}
	userdao := repository.Userdao
	usertest,err := userdao.QueryUserById(1)
	if err!=nil{
		fmt.Println("查询失败",err)
		return 
	}
	fmt.Println(usertest)
	return

}
func TestQueryUserByName(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return 
	}
	userdao := repository.Userdao
	usertest,err := userdao.QueryUserByName("lrs")
	if err!=nil{
		fmt.Println("查询失败",err)
		return 
	}
	fmt.Println(usertest)
	return

}


