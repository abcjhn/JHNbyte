package controller

import (

	"mydousheng/middleware"
	usersvc "mydousheng/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type UserLoginResponse struct{
	Response
	UserId int `json:"user_id,omitempty"`
	Token string `json:"token"`
}

type UserInfoResponse struct{
	Response
	User *usersvc.UserInfo `json:"user"`
}

func Register(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	userid,err := usersvc.Register(username,password)
	if err!= nil{
		c.JSON(http.StatusOK,UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}
	tokenstring,err := middleware.GenToken(userid)
	if err!= nil{
		 c.JSON(http.StatusOK,UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		 })
	}
	c.JSON(http.StatusOK,UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg: "注册成功",
		},
		UserId:userid ,
		Token: tokenstring,
	})

}
func Login(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	userid,err := usersvc.Login(username,password)
	if err!=nil{
		c.JSON(http.StatusOK,UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})

	}

	tokenstring,err := middleware.GenToken(userid)
	if err!=nil{
		c.JSON(http.StatusOK,UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})

	}
	c.JSON(http.StatusOK,UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg: "登录成功",
		},

		UserId: userid,
		Token: tokenstring,
	})

}

func UserInfo(c *gin.Context)  {
	useridstring := c.Query("user_id")
	userid,err := strconv.Atoi(useridstring)
	if err != nil{
		c.JSON(http.StatusOK,UserInfoResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}
	userinfo,err := usersvc.QueryUserInfo(userid)
	if err!=nil{
		c.JSON(http.StatusOK,UserInfoResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}
	c.JSON(http.StatusOK,UserInfoResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg: "查询到了呢",
		},
		User: userinfo,
	})

	
}