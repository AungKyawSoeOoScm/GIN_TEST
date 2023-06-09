package router

import (
	"gin_test/bulletin_board/controller"
	"gin_test/bulletin_board/interfaces"
	middlewares "gin_test/bulletin_board/middleware"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController, userController *controller.UsersController, postController *controller.PostController, usersInterface interfaces.UsersInterface) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static/")
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	apiRouter := router.Group("/")
	AuthRouter(apiRouter, authController)
	UsersRouter(apiRouter, usersInterface, userController)
	TagsRouter(apiRouter, postController, usersInterface)
	AuthForm(apiRouter, authController)
	return router
}

// Auth Form Routes
func AuthForm(router *gin.RouterGroup, authController *controller.AuthController) {
	authForm := router.Group("")
	{
		authForm.GET("/register", authController.RegisterForm)
		authForm.GET("/login", authController.LoginForm)
	}
}

// Auth Router
func AuthRouter(router *gin.RouterGroup, authController *controller.AuthController) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/logout", authController.Logout)
	}
}

// User Router
func UsersRouter(router *gin.RouterGroup, usersInterface interfaces.UsersInterface, usersController *controller.UsersController) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", middlewares.IsAuth(usersInterface), usersController.GetUsers)
	}
}

// PostRouter
func TagsRouter(router *gin.RouterGroup, PostsController *controller.PostController, userInterface interfaces.UsersInterface) {
	tagRouter := router.Group("/tags")
	{
		tagRouter.GET("/create", PostsController.CreateForm)
		// tagRouter.GET("/update/:tagId", PostsController.UpdateForm)
		tagRouter.GET("", middlewares.IsAuth(userInterface), PostsController.FindAll)
		// tagRouter.GET("/:tagId", middlewares.IsAuth(userInterface), PostsController.FindById)
		// tagRouter.POST("", middlewares.IsAuth(userInterface), func(ctx *gin.Context) {
		// 	userId := ctx.GetInt("Id")
		// 	PostsController.Create(ctx, userId)
		// })
		// tagRouter.POST("/:tagId", middlewares.IsAuth(userInterface), PostsController.Update)
		tagRouter.DELETE("/:tagId", middlewares.IsAuth(userInterface), PostsController.Delete)
	}
}
