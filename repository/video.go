package repository

import (
	"sync"
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


var videoDao *VideoDao
var videoOnce sync.Once


func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func ( *VideoDao)CreateVideo(video *Video) error{
	if err := GetDb().Create(video).Error; err!=nil{
		return err
	}
	return nil

}

func ( *VideoDao)QueryVideoList(num int) (*[]Video, error) {
	var videolist []Video
	if err:=GetDb().Limit(num).Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil
}
func ( *VideoDao)QueryVideoById(id int) (*Video,error) {
	var video Video
	if err:=GetDb().Where("id=?",id).Find(&video).Error; err!=nil{
		return nil,err
	}
	return &video,nil
}
func ( *VideoDao)QueryVideoByUid(uid int) (*[]Video,error){
	var videolist []Video
	if err := GetDb().Where("uid=?",uid).Find(&videolist).Error; err!=nil{
		return nil,err
	}
	return &videolist,nil
}

func ( *VideoDao)QueryVideoByIdList(id []int)(*[]Video,error)  {
	var videolist []Video
	if err := GetDb().Where("id in ?",id).Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil	
}

func( *VideoDao)QueryVideoListByTime(lasttime time.Time, len int)(*[]Video,error){
	var videolist []Video
	if err := GetDb().Where("Update_time < ?",lasttime).Order("update_time desc").Limit(len).Find(&videolist).Error;err!=nil{
		return nil,err
	}
	return &videolist,nil
}

func ( *VideoDao)IncFavoriteCount(id int) error{
	if err:=GetDb().Model(&Video{}).Where("id = ?",id).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func ( *VideoDao)DecFavoriteCount(id int) error{
	if err:=GetDb().Model(&Video{}).Where("id = ?",id).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func ( *VideoDao)IncCommentCount(id int) error{
	if err:=GetDb().Model(&Video{}).Where("id = ?",id).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}
func ( *VideoDao)DecCommentCount(id int) error{
	if err:=GetDb().Model(&Video{}).Where("id = ?",id).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error;err!=nil{
		return err
	}
	return nil
}