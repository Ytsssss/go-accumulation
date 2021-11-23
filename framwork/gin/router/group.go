package router

import (
	"github.com/Ytsssss/go_accumulation/framwork/gin/handler"
	"github.com/Ytsssss/go_accumulation/framwork/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()
	engine.Use(middleware.Log())
	group := engine.Group("/rest/v1")
	{
		group.POST("/upload", handler.Upload)
		group.POST("/upload/img", handler.UploadPic)
		group.POST("/upload/multi", handler.UploadMulti)
	}
	{
		group.GET("/query", handler.Query)
		group.GET("/query/:name/:age", handler.QueryParam)
		group.POST("/json", handler.Json)
	}

	if err := engine.Run(":8080"); err != nil {
		panic("error")
	}
}
