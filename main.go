package main

import (
	"mydousheng/repository"
	"os"
)

func main() {
	err := repository.InitDb()
	if err!=nil{
		os.Exit(-1)
	}
	db := repository.GetDb()
	DB,err := db.DB()
	if err!=nil{
		os.Exit(-1)
	}
	defer DB.Close()
	InitRouter()
}