package service

import (
	dao "gin_test/bulletin_board/dao/post"
	"gin_test/bulletin_board/data/response"

	"github.com/go-playground/validator/v10"
)

type PostsServiceImpl struct {
	PostDaoInterface dao.PostDaoInterface
	validate         *validator.Validate
}

func NewPostsRepositoryImpl(PostDaoInterface dao.PostDaoInterface, validate *validator.Validate) PostsService {
	return &PostsServiceImpl{
		PostDaoInterface: PostDaoInterface,
		validate:         validate,
	}
}

// Create implements TagsService
// func (t *PostsServiceImpl) Create(tags request.CreatePostsRequest, userId int) error {
// 	err := t.validate.Struct(tags)
// 	if err != nil {
// 		if validationErrs, ok := err.(validator.ValidationErrors); ok {
// 			return validationErrs
// 		}
// 		return err
// 	}

// 	tagModel := model.Posts{
// 		Title:        tags.Title,
// 		Description:  tags.Description,
// 		Status:       1,
// 		CreateUserId: userId,
// 	}
// 	err = t.postsInterface.Save(tagModel)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Delete implements TagsService
func (t *PostsServiceImpl) Delete(tagsId int) {
	t.PostDaoInterface.Delete(tagsId)
}

// FindAll implements TagsService
func (t *PostsServiceImpl) FindAll() []response.PostResponse {
	result := t.PostDaoInterface.FindAll()
	var tags []response.PostResponse
	for _, value := range result {
		tag := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			Status:      value.Status,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService
func (t *PostsServiceImpl) FindById(tagsId int) response.PostResponse {
	tagData := t.PostDaoInterface.FindById(tagsId)
	tagResponse := response.PostResponse{
		Id:          tagData.Id,
		Title:       tagData.Title,
		Description: tagData.Description,
		Status:      tagData.Status,
	}
	return tagResponse
}

// // Update implements TagsService
// func (t *PostsServiceImpl) Update(posts request.UpdatePostsRequest) {
// 	postData, err := t.postsInterface.FindById(posts.Id)
// 	helper.ErrorPanic(err)
// 	postData.Title = posts.Title
// 	postData.Description = posts.Description
// 	postData.Status = posts.Status
// 	t.postsInterface.Update(postData)
// }
