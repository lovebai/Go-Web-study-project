/*
 * @Descripttion: vue前端项目
 * @version: 1.0.0
 * @Author: Baishaodong
 * @Date: 2022-01-25 16:01:08
 * @LastEditors: Baishaodong
 * @LastEditTime: 2022-01-25 23:08:27
 * @BlogSite: https://www.xiaobaibk.com
 */
package controller

import (
	"fmt"
	"go-blog/dao"
	"go-blog/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

//注册
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Printf("注册的用户：%v 密码：%v\n", username, password)

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
	c.Redirect(301, "/")

}

//登录
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Printf("登录的用户：%v 密码：%v\n", username, password)

	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在")
		fmt.Printf("用户名不存在")

	} else {
		if u.Password != password {
			fmt.Print("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Printf("登录成功")
			c.Redirect(301, "/")
		}
	}

}

//首页
func ToIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//注册
func ToReg(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

//登录
func ToLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//博客操作
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	// s := c.Query("id")
	s := c.Param("pid")
	fmt.Printf("pid=%v", s)
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}

func PostDetails(c *gin.Context) {
	// s := c.Query("id")
	s := c.Param("title")
	fmt.Printf("title=%v", s)
	p := dao.Mgr.GetpostTitle(s)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
