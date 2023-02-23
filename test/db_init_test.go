package test

import (
	"mydousheng/repository"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	type UserTest struct {
		gorm.Model
		Name     string
		Birthday time.Time
	}
	err := repository.Init();
	if err!=nil{
		os.Exit(-1)
	}
	db := repository.Db
	db.AutoMigrate(&UserTest{})
	return

}
