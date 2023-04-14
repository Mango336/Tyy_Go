package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpLoadFile 上传文件
func UpLoadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	dst := "./files/" + file.Filename
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("%s is uploaded...", file.Filename))
}

// UpLoadFiles 上传多个文件
func UpLoadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		log.Println(file.Filename)
		dst := "./files/" + file.Filename
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files are uploaded...", len(files)))
}
