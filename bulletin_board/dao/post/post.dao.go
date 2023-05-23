package postDao

import (
	"fmt"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"

	"gorm.io/gorm"
)

type PostDao struct {
	DB *gorm.DB
}

func (postDao *PostDao) Delete(postId int) {
	var post model.Posts
	post.Id = postId
	postDao.DB.Delete(&post)
}

func (postDao *PostDao) FindAll() []model.Posts {
	var posts []model.Posts
	result := postDao.DB.Find(&posts)
	helper.ErrorPanic(result.Error)
	fmt.Print(posts)
	return posts
}
func (postDao *PostDao) FindById(postId int) model.Posts {
	var post model.Posts
	result := postDao.DB.First(&post, postId)
	helper.ErrorPanic(result.Error)
	return post
}

func NewPostDao(DB *gorm.DB) PostDaoInterface {
	return &PostDao{DB: DB}
}
