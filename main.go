package main

import (
	"go-lang/controllers"
	"go-lang/initializers"

	// "go-lang/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	// migrate.Migrate()
}

func main() {
	r := gin.Default()
	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/page/:limit/:page", controllers.PaginatePostsByPage)
		postRoutes.GET("/offset/:limit/:offset", controllers.PaginatePostsByOffset)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
	r.Run()
}
