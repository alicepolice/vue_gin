package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/structs"
	"wolflong.com/vue_gin/variable"
)

func QueryGameList(c *gin.Context) {
	db := variable.DB
	rows, err := db.Query(`select a.title,a.text,img,price,web,if(group_concat(c.title separator "#") is null ,"", group_concat(c.title separator "#")) as tag from game a left join gametag b on a.id = b.gameid left join tag c on b.tagid = c.id group by a.id;`)
	checkError(err)
	defer rows.Close()
	var GameList []structs.Gamelist
	for rows.Next() {
		var g structs.Gamelist
		var tag string
		err = rows.Scan(&g.Title, &g.Text, &g.Img, &g.Price, &g.Web, &tag)
		g.Tag = strings.Split(tag, "#") // 这里暂且由服务端完成分解
		checkError(err)
		GameList = append(GameList, g)
	}
	c.JSON(200, GameList)
}
