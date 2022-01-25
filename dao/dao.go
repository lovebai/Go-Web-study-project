/*
 * @Descripttion: vue前端项目
 * @version: 1.0.0
 * @Author: Baishaodong
 * @Date: 2022-01-25 16:01:55
 * @LastEditors: Baishaodong
 * @LastEditTime: 2022-01-25 23:36:14
 * @BlogSite: https://www.xiaobaibk.com
 */
package dao

import (
	"fmt"
	"go-blog/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User
	// 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
	GetpostTitle(title string) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	//mysql
	// dsn := "root:root@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//sqlite
	fmt.Println("创建数据库连接中...")
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(&model.User{})
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user

}

// 博客操作
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	fmt.Printf("pid=%v\n", pid)
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
func (mgr *manager) GetpostTitle(title string) model.Post {
	fmt.Printf("pid=%v\n", title)
	var post model.Post
	mgr.db.First(&post, "title=?", title)
	return post
}
