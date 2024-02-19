package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (u *UserHandler) RegisterUserRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	{
		ug.GET("list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "all user list",
			})
		})
		ug.POST("signup", func(c *gin.Context) {
			type signupReq struct {
				Uname string `json:"uname"`
				Email string `json:"email"`
				Cred  string `json:"cred"`
			}
			var req signupReq
			//Bind方法会根据Content-Type 解析请求数据
			if err := c.Bind(&req); err != nil {
				fmt.Printf("error %v \n", err)
				return
			}
			fmt.Printf("req %v \n", req)
			c.JSON(http.StatusOK, gin.H{
				"message": req.Uname + " signup successful",
			})
		})
		ug.GET("find/:name", func(ctx *gin.Context) {
			//参数路由 /user/find/{name}
			name := ctx.Param("name")
			fmt.Printf("%s user found", name)
			ctx.JSON(http.StatusOK, gin.H{
				"message": name + " found",
			})
		})
		ug.DELETE("unsubscribe", func(c *gin.Context) {
			//通过路由参数查找 /user/unsubscribe?id=123
			uid := c.Query("id")
			c.JSON(http.StatusOK, gin.H{
				"message": uid + " user unsubscribed",
			})
		})
	}
	ug.POST("signin", func(c *gin.Context) {
		type signinReq struct {
			Uname string `json:"uname"`
			Cred  string `json:"cred"`
		}
		var req signinReq
		if err := c.Bind(&req); err != nil {
			fmt.Printf("error parsing request %v \v", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": req.Uname + " have signin",
		})
	})
}
