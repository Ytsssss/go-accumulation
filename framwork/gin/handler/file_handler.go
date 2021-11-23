package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusInternalServerError, "upload error")
	}

	if err = context.SaveUploadedFile(file, file.Filename); err != nil {
		context.String(http.StatusInternalServerError, "save upload file error")
	}
	context.String(http.StatusOK, "%s upload success", file.Filename)
}

func UploadPic(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}
	if header.Size > 1024*1024*2 {
		fmt.Println("文件太大了")
		return
	}
	if header.Header.Get("Content-Type") != "image/png" {
		fmt.Println("只允许上传png图片")
		return
	}

	if err = ctx.SaveUploadedFile(header, "img/"+header.Filename); err != nil {
		ctx.String(http.StatusInternalServerError, "save upload file error")
	}
	ctx.String(http.StatusOK, header.Filename)
}

func UploadMulti(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "get multipart form error")
	}
	files := form.File["files"]

	for _, file := range files {
		if err = ctx.SaveUploadedFile(file, "img/"+file.Filename); err != nil {
			ctx.String(http.StatusInternalServerError, "save upload file error")
		}
	}
	ctx.String(http.StatusOK, "success upload files %d", len(files))
}
