package main

import (
	"github.com/gin-gonic/gin"
	"wolflong.com/vue_gin/router"
	"wolflong.com/vue_gin/utils"
)

func init() {
	utils.MySqlDB()
}

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":1314")
}
