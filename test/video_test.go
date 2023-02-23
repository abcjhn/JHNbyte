package test

import (
	"fmt"
	"mydousheng/repository"
	"testing"
)
func TestCreateVideo(t *testing.T) {

	if err:=repository.Init();err!=nil{
		return 
	}
	db :=repository.Db
	db.Migrator().CreateTable(&repository.Video{})
	videodao := repository.Videodao
	videodao.CreateVideo(&repository.Video{Tittle: "bear",Uid: 5,CommentCount: 50,FavoriteCount: 50})
}
func TestQueryVideoList(t *testing.T)  {
	if err:=repository.Init();err!=nil{
		return 
	}
	videodao := repository.Videodao
	videolist,err := videodao.QueryVideoList(2)
	if err != nil{
		fmt.Println("拉取视频列表失败",err)
		return 
	}
	fmt.Println(videolist)
	return 
}
func TestQueryVideoById(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao := repository.Videodao
	video,err:=videodao.QueryVideoById(1)
	if err!=nil{
		fmt.Println("根据id拉取视频失败")
		return
	}
	fmt.Println(video)
	return
}
func TestQueryVideoByUid(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao := repository.Videodao
	videolist,err:=videodao.QueryVideoByUid(3)
	if err!=nil{
		fmt.Println("根据uid拉取视频失败")
		return
	}
	fmt.Println(videolist)
	return

}
func TestQueryVideoByIdList(t *testing.T) {

	if err:=repository.Init();err!=nil{
		return
	}
	videodao := repository.Videodao
	videolist,err:=videodao.QueryVideoByIdList([]int{1,2,3})
	if err!=nil{
		fmt.Println("根据id列表拉取视频失败")
		return
	}
	fmt.Println(videolist)
	return

}

func TestIncFavoriteCount(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao:=repository.Videodao
	if err:= videodao.IncFavoriteCount(9);err!=nil{
		fmt.Println("增加关注数失败")
		return
	}
	return

}
func TestDecFavoriteCount(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao:=repository.Videodao
	if err:= videodao.DecFavoriteCount(9);err!=nil{
		fmt.Println("取关失败")
		return
	}
	return

}
func TestIncCommentCount(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao:=repository.Videodao
	if err:= videodao.IncCommentCount(9);err!=nil{
		fmt.Println("添加评论失败")
		return
	}
	return

}
func TestDecCommentCount(t *testing.T) {
	if err:=repository.Init();err!=nil{
		return
	}
	videodao:=repository.Videodao
	if err:= videodao.DecCommentCount(9);err!=nil{
		fmt.Println("撤销评论失败")
		return
	}
	return

}