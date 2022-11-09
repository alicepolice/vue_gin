package router

import (
	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/controller"
)

func Router(r *gin.Engine) {
	r.GET("/", controller.HelloWorld)
	r.GET("/queryGameblog", controller.QueryGameBlog)
	r.GET("/queryGamelist", controller.QueryGameList)

	comment := r.Group("/comment")
	{
		comment.GET("/query", controller.QueryComment)
		comment.POST("/delete", controller.DeleteComment)
		comment.POST("/insert", controller.InsertComment)
	}

	post := r.Group("/post")
	{
		post.GET("/query", controller.QueryPost)
	}
}
