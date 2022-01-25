/*
 * @Descripttion: vue前端项目
 * @version: 1.0.0
 * @Author: Baishaodong
 * @Date: 2022-01-25 16:00:51
 * @LastEditors: Baishaodong
 * @LastEditTime: 2022-01-25 23:09:17
 * @BlogSite: https://www.xiaobaibk.com
 */
package router

import (
	"go-blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.ToIndex)
	e.GET("/register", controller.ToReg)
	e.POST("/register", controller.AddUser)
	e.GET("/login", controller.ToLogin)
	e.POST("/login", controller.Login)
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	// e.GET("/post/:pid", controller.PostDetail)
	e.GET("/post/:title", controller.PostDetails)
	e.Run()
}
