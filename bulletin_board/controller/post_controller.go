package controller

import (
	"fmt"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/post"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	tagsService service.PostsService
}

func NewPostsController(service service.PostsService) *PostController {
	return &PostController{
		tagsService: service,
	}
}

// // create controller
// func (controller *PostController) Create(ctx *gin.Context, userId int) {
// 	title := ctx.PostForm("title")
// 	description := ctx.PostForm("description")
// 	createTagsRequest := request.CreatePostsRequest{
// 		Title:       title,
// 		Description: description,
// 		UserId:      userId,
// 	}
// 	fmt.Print(userId)
// 	err := controller.tagsService.Create(createTagsRequest, userId)
// 	if err != nil {
// 		if validationErr, ok := err.(validator.ValidationErrors); ok {
// 			errorMessages := make(map[string]string)
// 			for _, fieldErr := range validationErr {
// 				fieldName := fieldErr.Field()
// 				errorMessage := ""
// 				switch fieldErr.Tag() {
// 				case "required":
// 					errorMessage = "Field is required"
// 				case "min":
// 					errorMessage = fmt.Sprintf("Field must be at least %s characters long", fieldErr.Param())
// 				case "max":
// 					errorMessage = fmt.Sprintf("Field must not exceed %s characters", fieldErr.Param())
// 				default:
// 					errorMessage = "Field validation failed"
// 				}
// 				errorMessages[fieldName] = errorMessage

// 			}
// 			fmt.Println(errorMessages)
// 			ctx.HTML(http.StatusBadRequest, "create.html", gin.H{
// 				"Errors": errorMessages,
// 			})
// 			return
// 		}
// 	} else {
// 		ctx.Redirect(http.StatusFound, "/tags")
// 	}
// }

// // update controller
// func (controller *PostController) Update(ctx *gin.Context) {
// 	tagId := ctx.Param("tagId")
// 	title := ctx.PostForm("title")
// 	description := ctx.PostForm("description")
// 	statusValue := ctx.PostForm("status")
// 	fmt.Println(statusValue)
// 	fmt.Println(tagId)
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)
// 	if method := ctx.Request.Header.Get("X-HTTP-Method-Override"); method == "PUT" {
// 		ctx.Request.Method = "PUT"
// 	}
// 	updateTagsRequest := request.UpdatePostsRequest{
// 		Id:          id,
// 		Title:       title,
// 		Description: description,
// 		Status:      2,
// 	}
// 	if statusValue == "on" {
// 		updateTagsRequest.Status = 1
// 	}
// 	fmt.Println(updateTagsRequest)
// 	if err := ctx.ShouldBind(&updateTagsRequest); err != nil {
// 		ctx.HTML(http.StatusBadRequest, "update.html", gin.H{
// 			"Tag":    updateTagsRequest,
// 			"Errors": err.Error(),
// 		})
// 		return
// 	}
// 	controller.tagsService.Update(updateTagsRequest)
// 	ctx.Redirect(http.StatusFound, "/tags")
// }

// delete controller
func (controller *PostController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	// Check for method override header
	controller.tagsService.Delete(id)
	ctx.Redirect(http.StatusFound, "/tags")
}

// findById controller
// func (controller *PostController) FindById(ctx *gin.Context) {
// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)
// 	tagsResponse := controller.tagsService.FindById(id)
// 	webResponse := response.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   tagsResponse,
// 	}
// 	ctx.Header("Content-Type", "application/json")
// 	ctx.JSON(http.StatusOK, webResponse)
// }

// findAll controller
func (controller *PostController) FindAll(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("token")
	fmt.Println(cookie)
	if err != nil || cookie.Value == "" {
		fmt.Print("No token")
		return
	}
	tagResponse := controller.tagsService.FindAll()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"tags": tagResponse,
	})
}

func (controller *PostController) CreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.html", gin.H{})
}

// func (controller *PostController) UpdateForm(ctx *gin.Context) {
// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)
// 	tag := controller.tagsService.FindById(id)
// 	ctx.HTML(http.StatusOK, "update.html", gin.H{
// 		"Tag": tag,
// 	})
// 	fmt.Print(tag)
// }
