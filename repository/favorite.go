package repository

import "time"
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
var Favdao *FavoriteDao
func (favdao *FavoriteDao)CreateFavorite(fav *Favorite) error{
	if err:=Db.Create(&fav).Error;err!=nil{
		return err
	}
	return nil
}
func (favdao *FavoriteDao)QueryByUidandVid(uid int,vid int)(bool,error){
	var fav Favorite
	if err:=Db.Model(&fav).Where("uid=?",uid).Where("vid=?",vid).Error;err!=nil{
		return false,err
	}
	return fav.IsFavorite,nil
}
func (favdao *FavoriteDao)QueryFavlistByUid(uid int) (*[]Favorite,error) {
	var favlist []Favorite
	if err:=Db.Where("uid=?",uid).Find(&favlist).Error;err!=nil{
		return nil,err
	}
	return &favlist,nil
}
func (favdao *FavoriteDao)UpdateIsFavorite(uid int,vid int,IsFavorite bool) error  {
	var fav *Favorite
	if err:= Db.Model(&fav).Where("uid=?",uid).Where("vid=?",vid).UpdateColumn("is_favorite",IsFavorite).Error;err!=nil{
		return err
	}
	return nil
	
}