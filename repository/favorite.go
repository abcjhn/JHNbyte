package repository

import (
	"sync"
	"time"
)
type Favorite struct{
	Id 			int				`gorm"column:id"`	
	uid 		int				`gorm:"column:uid"`
	vid 		int				`gorm:"column:vid"`
	IsFavorite 	bool			`gorm:"column:is_favorite"`
	CreatedAt 	time.Time		`gorm:"column:created_time"`
	UpdatedAt 	time.Time		`gorm:"column:updated_time"`
}

func (f *Favorite)TableName() string{
	return "favorite"
}
type FavoriteDao struct{}
var favDao *FavoriteDao
var favOnce sync.Once

func NewFavoriteDaoIntance()  *FavoriteDao{
	favOnce.Do(func() {
		favDao = &FavoriteDao{}
	})
	return favDao
	
}

func ( *FavoriteDao)CreateFavorite(fav *Favorite) error{
	if err:=GetDb().Create(&fav).Error;err!=nil{
		return err
	}
	return nil
}
func ( *FavoriteDao)QueryByUidandVid(uid int,vid int)(bool,error){
	var fav Favorite
	if err:=GetDb().Model(&fav).Where("uid=?",uid).Where("vid=?",vid).Error;err!=nil{
		return false,err
	}
	return fav.IsFavorite,nil
}
func ( *FavoriteDao)QueryFavlistByUid(uid int) (*[]Favorite,error) {
	var favlist []Favorite
	if err:=GetDb().Where("uid=?",uid).Find(&favlist).Error;err!=nil{
		return nil,err
	}
	return &favlist,nil
}
func ( *FavoriteDao)UpdateIsFavorite(uid int,vid int,IsFavorite bool) error  {
	var fav *Favorite
	if err:= GetDb().Model(&fav).Where("uid=?",uid).Where("vid=?",vid).UpdateColumn("is_favorite",IsFavorite).Error;err!=nil{
		return err
	}
	return nil
	
}