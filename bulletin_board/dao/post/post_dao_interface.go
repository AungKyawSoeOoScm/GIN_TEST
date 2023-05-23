package postDao

import "gin_test/bulletin_board/model"

type PostDaoInterface interface {
	FindAll() []model.Posts
	FindById(postId int) model.Posts
	// Create(post model.Posts)
	// Update(post model.Posts)
	Delete(postId int)
}
