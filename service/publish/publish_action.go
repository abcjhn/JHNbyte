package publishsvc

import (
	"mydousheng/repository"
)



func PublishVideo(myid int,filename string,covername string )error  {
	length := len(filename)

	video:=&repository.Video{
		Uid: myid,
		PlayUrl: filename,
		CoverUrl: covername,
		FavoriteCount: 0,
		CommentCount: 0,
		Tittle: filename[:length-4],
		IsDeleted: false,

	}
	if err:=repository.Videodao.CreateVideo(video);err!=nil{
		return err
	}
	return nil
}