package repository

import "time"
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
var Cmtdao *CommentDao

func(cmtdao *CommentDao) CreateCmt(cmt *Comment)error{
	if err:=Db.Create(&cmt).Error;err!=nil{
		return err
	}
	return nil
}

func(cmtdao *CommentDao) QueryCmtByVid(vid int) (*[]Comment,error) {
	var  cmtlist []Comment
	if err:=Db.Where("vid=?",vid).Find(&cmtlist).Error;err!=nil{
		return nil,err
	}
	return &cmtlist,nil
	
}