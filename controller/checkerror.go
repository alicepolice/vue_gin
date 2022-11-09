package controller

import "github.com/gin-gonic/gin"

func checkError(err error, handlers ...gin.HandlerFunc) {
	if err != nil {
		panic(err.Error())
	}
}
