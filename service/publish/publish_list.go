package publishsvc

import "mydousheng/repository"


type VideoInfo struct {
	Id            int        	`json:"id,omitempty"`                              		//视频唯一标识
	Author        *UserInfo    	`json:"author,omitempty"`                               //视频作者信息
	PlayUrl       string        `json:"play_url,omitempty"`            					//视频播放地址
	CoverUrl      string        `json:"cover_url,omitempty"`        					//视频封面地址
	FavoriteCount int 			`json:"favorite_count,omitempty"`	 					//视频的点赞总数
	CommentCount  int    		`json:"comment_count,omitempty"`						//视频的评论总数
	IsFavorite    bool         	`json:"is_favorite,omitempty"`							//true-已点赞 false-未点赞
	Title         string        `json:"title,omitempty"`                      			//视频标题
}

type UserInfo struct{
	Id int 			`json:"id,omitempty"`  
	name string		`json:"name,omitempty"`
}


func GetPublishList(uid int,myid int) ([]*VideoInfo,error){
	VideoPath:= "http://139.196.75.69/mov/"
	ImgPath :=  "http://139.196.75.69/pic/"
	videolist,err := repository.NewVideoDaoInstance().QueryVideoByUid(uid)
	if err!=nil{
		return nil,err
	}
	var videoinfolist []*VideoInfo
	for _,video := range *videolist{
		user,err := repository.NewUserDaoInstance().QueryUserById(uid)
		if err!=nil{
			return nil,err
		}
		userinfo := &UserInfo{
			Id: user.Id,
			name: user.Username,
		}
		isfavorite,err :=repository.NewFavoriteDaoIntance().QueryByUidandVid(video.Id,myid)
		if err!=nil{
			isfavorite=false
		}
		videoinfo := &VideoInfo{
			Id: video.Id,
			Author: userinfo,
			PlayUrl: VideoPath+ video.PlayUrl,
			CoverUrl: ImgPath+video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: isfavorite,
			Title: video.Tittle,
		}
		videoinfolist=append(videoinfolist, videoinfo)
	}
	return videoinfolist,err

}