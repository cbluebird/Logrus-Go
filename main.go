package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"logs/app/apiException"
	"logs/app/midwares"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	r.GET("/hello", func(c *gin.Context) {
		_ = c.AbortWithError(200, errors.Cause(e3))
		// c.JSON：返回JSON格式的数据
	})
	r.GET("/hello1", func(c *gin.Context) {
		_ = c.AbortWithError(200, apiException.HttpTimeout)
		// c.JSON：返回JSON格式的数据
	})
	r.GET("/hello2", func(c *gin.Context) {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		// c.JSON：返回JSON格式的数据
	})

	err := r.Run()
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
