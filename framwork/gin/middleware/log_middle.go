package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		log.Println("begin handle")
		context.Set("request", "middleware")
		context.Next()
		status := context.Writer.Status()
		log.Println("finish handle", status)
		log.Println("cost time: ", time.Since(t))
	}
}
