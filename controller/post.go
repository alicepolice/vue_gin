package controller

import (
	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/structs"
	"wolflong.com/vue_gin/variable"
)

func QueryPost(c *gin.Context) {
	db := variable.DB
	pid := c.Query("pid")
	rows, err := db.Query(`select id,bgcolor,textcolor,headimg,videosrc,imgs,html from posts where id = ?`, pid)
	checkError(err)
	defer rows.Close()
	var Post []structs.Post
	rows.Next()
	var g structs.Post
	err = rows.Scan(&g.ID, &g.Bgcolor, &g.Textcolor, &g.Headimg, &g.Videosrc, &g.Imgs, &g.Html)
	checkError(err)
	Post = append(Post, g)
	c.JSON(200, Post)
}
