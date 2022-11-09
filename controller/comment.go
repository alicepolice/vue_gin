package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/structs"
	"wolflong.com/vue_gin/variable"
)

func QueryComment(c *gin.Context) {
	db := variable.DB
	pid := c.Query("pid")
	rows, err := db.Query(`select id,uid,text,pid,date,name from comments join users using(uid) where pid = ?`, pid)
	checkError(err)
	defer rows.Close()
	var res []structs.Comment
	for rows.Next() {
		var c structs.Comment
		err = rows.Scan(&c.ID, &c.UID, &c.Text, &c.Pid, &c.Date, &c.Name)
		checkError(err)
		res = append(res, c)
	}
	c.JSON(200, res)
}

func DeleteComment(c *gin.Context) {
	db := variable.DB
	cid := c.PostForm("id")
	res, err := db.Exec("delete from comments where id = ?", cid)
	checkError(err)
	n, err := res.RowsAffected()
	checkError(err)
	if n == 0 {
		c.JSON(501, gin.H{
			"message": "failure",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func InsertComment(c *gin.Context) {
	db := variable.DB
	uid := c.PostForm("uid")
	pid := c.PostForm("pid")
	text := c.PostForm("text")
	res, err := db.Exec(`INSERT INTO comments(uid,text,pid,date) values(?,?,?,?)`,
		uid, text, pid, time.Now().UnixMilli())
	checkError(err)
	n, err := res.RowsAffected()
	checkError(err)
	if n == 0 {
		c.JSON(501, gin.H{
			"message": "failure",
		})
		c.Abort()
		return
	}
	n, err = res.LastInsertId()
	checkError(err)
	rows, err := db.Query(`select id,uid,text,pid,date,name from comments join users using(uid) where id = ?`, n)
	checkError(err)
	defer rows.Close()
	rows.Next()
	var cm structs.Comment
	rows.Scan(&cm.ID, &cm.UID, &cm.Text, &cm.Pid, &cm.Date, &cm.Name)
	c.JSON(200, cm)
}
