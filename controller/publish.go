package controller

import (
	"bytes"
	"fmt"
	"image"
	"log"
	publishsvc "mydousheng/service/publish"
	"net/http"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)


type VideoListResponse struct{
	Response
	VideoList []*publishsvc.VideoInfo `json:"video_list"`
}


func PublishList(c *gin.Context)  {
	uid,_ := c.Get("uid")
	myuid := uid.(int)
	queryuidstring := c.Query("user_id")
	queryuid,_ := strconv.Atoi(queryuidstring)
	videolist,err := publishsvc.GetPublishList(queryuid,myuid)
	if err!=nil{
		c.JSON(http.StatusOK,VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}
	c.JSON(http.StatusOK,VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg: "查询成功",
		},
		VideoList: videolist,
	})

	
}

func PublishAction(c *gin.Context)  {
	VideoPath:= "/var/www/html/mov/"
	uidget,_ := c.Get("uid")
	uid,_ := uidget.(int)
	tittle := c.PostForm("title")
	file,err := c.FormFile("data")
	if err!=nil{
		c.JSON(http.StatusOK,Response{
			StatusCode: 1,
			StatusMsg: err.Error(),
		})

	}
	log.Println(file.Filename)

	//解析PlayUrl和CoverUrl
	length := len(file.Filename)
	PlayUrl := tittle+file.Filename[length-4:]
	videodst := VideoPath+PlayUrl
	c.SaveUploadedFile(file,videodst)

	img,err := GetSnapshot(videodst)
	if err!=nil{
		c.JSON(http.StatusOK,Response{
			StatusCode: 1,
			StatusMsg: err.Error(),
		})

	}

	CoverUrl,err := SaveImg(img,tittle)
	if err!=nil{
		c.JSON(http.StatusOK,Response{
			StatusCode: 1,
			StatusMsg: err.Error(),
		})

	}



	if err:=publishsvc.PublishVideo(uid,PlayUrl,CoverUrl);err!=nil{
		c.JSON(http.StatusOK,Response{
			StatusCode: 1,
			StatusMsg: err.Error(),
		})

	}
	c.JSON(http.StatusOK,Response{
		StatusCode: 0,
		StatusMsg: "上传成功",
	})


	
}

func GetSnapshot(videoPath string) (image.Image,  error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
	   Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
	   WithOutput(buf, os.Stdout).Run()
 
	if err != nil {
	   log.Fatal("生成缩略图失败：", err)
	   return nil,err
	}
 
	img, err := imaging.Decode(buf)
	if err != nil {
	   log.Fatal("生成缩略图失败：", err)
	   return nil, err
	}
 
	return img, nil
 }
 

func SaveImg(img image.Image,imgfinalpath string) (string,error) {
	ImgPath :=  "/var/www/html/pic/"
	CoverUrl := imgfinalpath + ".png"
	snapshotPath := ImgPath+CoverUrl
	err := imaging.Save(img, snapshotPath+".png")
	if err != nil {
	   log.Fatal("生成缩略图失败：", err)
	   return "", err
	}
	return CoverUrl,nil

}

   

