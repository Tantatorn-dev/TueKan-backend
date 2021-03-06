package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

//Post routes for Post API
func Post(app *echo.Echo) {
	postController := controller.NewPostController(db.DB)

	postRoute := app.Group("/post")
	postRoute.GET("/", postController.GetPostList)
	postRoute.GET("/posting/:id", postController.GetPosting)
	postRoute.POST("/", postController.CreatePost)

}
