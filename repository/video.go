package repository

import (
	"time"

	"gorm.io/gorm"
)

type Video struct{
	Id				int			`gorm:"column:id"`
	Uid				int			`gorm:"column:uid" `
	PlayUrl			string		`gorm:"column:paly_url"`
	CoverUrl 		string		`gorm:"column:cover_url"`
	FavoriteCount 	int			`gorm:"column:favorite_count"`
	CommentCount 	int			`gorm:"column:comment_count"`
	Tittle 			string		`gorm:"column:tittle"`
	CreatedAt		time.Time	`gorm:"column:created_time"`
	UpdatedAt		time.Time	`gorm:"column:updated_time"`
	IsDeleted		bool		`gorm:"column:is_deleted"`
}
func (v *Video)TableName()	string  {
	return "video"
}

type VideoDao struct{}

var Videodao *VideoDao
func (videodao *VideoDao)CreateVideo(video *Video) error{
	if err := Db.Create(video).Error; err!=nil{
		return err
	}
	return nil

}

func (videodao *VideoDao)QueryVideoList(num int) (*[]Video, error) {
	var videolist []Video
	if err:=Db.Limit(num).Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil
}
func (videodao *VideoDao)QueryVideoById(id int) (*Video,error) {
	var video Video
	if err:=Db.Where("id=?",id).Find(&video).Error; err!=nil{
		return nil,err
	}
	return &video,nil
}
func (videodao *VideoDao)QueryVideoByUid(uid int) (*[]Video,error){
	var videolist []Video
	if err := Db.Where("uid=?",uid).Find(&videolist).Error; err!=nil{
		return nil,err
	}
	return &videolist,nil
}

func (videodao *VideoDao)QueryVideoByIdList(id []int)(*[]Video,error)  {
	var videolist []Video
	if err := Db.Where("id in ?",id).Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil	
}

func(videodao *VideoDao)QueryVideoListByTime(lasttime time.Time)(*[]Video,error){
	var videolist []Video
	if err := Db.Where("Update_time < ?",lasttime).Order("update_time desc").Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil
}

func (videodao *VideoDao)IncFavoriteCount(id int) error{
	if err:=Db.Model(&Video{}).Where("id = ?",id).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func (videodao *VideoDao)DecFavoriteCount(id int) error{
	if err:=Db.Model(&Video{}).Where("id = ?",id).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func (videodao *VideoDao)IncCommentCount(id int) error{
	if err:=Db.Model(&Video{}).Where("id = ?",id).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func (videodao *VideoDao)DecCommentCount(id int) error{
	if err:=Db.Model(&Video{}).Where("id = ?",id).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}