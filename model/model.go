/*
 * @Descripttion: vue前端项目
 * @version: 1.0.0
 * @Author: Baishaodong
 * @Date: 2022-01-25 16:01:33
 * @LastEditors: Baishaodong
 * @LastEditTime: 2022-01-25 22:18:22
 * @BlogSite: https://www.xiaobaibk.com
 */
package model

import "github.com/jinzhu/gorm"

//用户模型
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

//文章模型
type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
