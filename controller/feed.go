package controller

import (
	feedsvc "mydousheng/service/feed"
	"net/http"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct{
	Response
	NextTime time.Time 				`json:"next_time,omitempty"`
	VideoList []*feedsvc.VideoInfo	`json:"video_list,omitempty"`
}

func Feed(c *gin.Context) {
	myuidstring,_ := c.Get("uid")
	myuid := myuidstring.(int)
	CurrentTime := time.Now().Unix()
	CurrentTimeString := strconv.FormatInt(CurrentTime,10)
	lasttimestr := c.DefaultQuery("latest_time",CurrentTimeString)
	lasttimeint,err := strconv.ParseInt(lasttimestr,10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: "时间请求解析失败",
			},
		})
	}
	LastTime := time.Unix(lasttimeint,0)
	videolist,nexttime,err := feedsvc.Feed(LastTime,10,myuid) 
	if err!=nil{
		c.JSON(http.StatusOK,FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
			NextTime: nexttime,
			VideoList: videolist,
		})
	}
}