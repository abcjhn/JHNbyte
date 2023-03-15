package repository

import (
	"sync"
	"time"
)
type Comment struct{
	Id			int				`gorm:"column:id"`
	Uid			int				`gorm:"column:uid"`
	Vide 		int				`gorm:"column:vid"`
	Content		string			`gorm:"column:content"`
	IsDeleted 	bool			`gorm:"column:is_deleted"`
	CreatedAt 	time.Time		`gorm:"column:created_time"`
	UpdatedAt 	time.Time		`gorm:"column:updated_time"`
}
type CommentDao struct{}
var cmtDao *CommentDao
var cmtOnce sync.Once

func NewCommentDaoIntance() *CommentDao{
	cmtOnce.Do(func() {
		cmtDao = &CommentDao{}
	})
	return cmtDao
}

func( *CommentDao) CreateCmt(cmt *Comment)error{
	if err:=GetDb().Create(&cmt).Error;err!=nil{
		return err
	}
	return nil
}

func( *CommentDao) QueryCmtByVid(vid int) (*[]Comment,error) {
	var  cmtlist []Comment
	if err:=GetDb().Where("vid=?",vid).Find(&cmtlist).Error;err!=nil{
		return nil,err
	}
	return &cmtlist,nil
	
}