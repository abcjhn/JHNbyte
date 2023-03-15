package feedsvc

import (

	"mydousheng/repository"
	"time"
)
type VideoInfo struct {
	Id            int                                            //视频唯一标识
	Author        *UserInfo                                     	//视频作者信息
	PlayUrl       string                    						//视频播放地址
	CoverUrl      string                							//视频封面地址
	FavoriteCount int 				 							//视频的点赞总数
	CommentCount  int    										//视频的评论总数
	IsFavorite    bool         										//true-已点赞 false-未点赞
	Title         string                              				//视频标题
}

type UserInfo struct{
	Id int
	name string
}

func Feed(lasttime time.Time,len int,myid int)(*[]VideoInfo,error)  {
	VideoPath:= "http://139.196.75.69/mov/"
	ImgPath :=  "http://139.196.75.69/pic/"
	videolist,err := repository.Videodao.QueryVideoListByTime(lasttime)
	if err != nil {
		return nil, err
	}
	 videoinfolist :=make([]VideoInfo,len)
	for _,video := range *videolist{
		user,err := repository.Userdao.QueryUserById(video.Uid)
		if err != nil {
			return nil, err
		}
		userinfo := &UserInfo{
			Id: user.Id,
			name: user.Username,
		}
		isfavorite,err := repository.Favdao.QueryByUidandVid(myid,video.Id)
		if err != nil {
			isfavorite = false
		}
		videoinfo := &VideoInfo{
			Id: video.Id,
			Author: userinfo,
			PlayUrl:VideoPath+video.PlayUrl ,
			CoverUrl:ImgPath+video.CoverUrl ,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: isfavorite,
			Title: video.Tittle,
		}
		videoinfolist = append(videoinfolist, *videoinfo)
	}
	
	return &videoinfolist,err
	
}