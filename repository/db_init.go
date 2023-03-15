package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// user     = debian-sys-maint
// password = SxnhK2jjmaDMIvwB
// port = 3306


var _db *gorm.DB
func InitDb() error {
	host := "127.0.0.1"
	username := "debian-sys-maint"
	password := "SxnhK2jjmaDMIvwB"
	port := 3306
	dbname := "dousheng"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	var err error

	_db,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	
	if err != nil{
		panic("连接数据库失败, error=" + err.Error())
	}
	_db.AutoMigrate(&User{})
	_db.AutoMigrate(&Video{})
	_db.AutoMigrate(&Favorite{})
	_db.AutoMigrate(&Comment{})
	
	return err
}

func GetDb() *gorm.DB{
	return _db
}