package controller

import (
	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/structs"
	"wolflong.com/vue_gin/variable"
)

func QueryGameBlog(c *gin.Context) {
	db := variable.DB
	rows, err := db.Query("SELECT title,text,img FROM gameblog order by id desc")
	checkError(err)
	defer rows.Close()
	var gameBlogs []structs.Gameblog
	for rows.Next() {
		var gameBlog structs.Gameblog
		err = rows.Scan(&gameBlog.Title, &gameBlog.Text, &gameBlog.Img)
		checkError(err)
		gameBlogs = append(gameBlogs, gameBlog)
	}
	c.JSON(200, gameBlogs)
}
