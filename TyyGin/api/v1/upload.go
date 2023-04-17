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
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println(err)
	} else {
		c.String(http.StatusOK, fmt.Sprintf("%s is uploaded...", file.Filename))
	}
}

// UpLoadFiles 上传多个文件
func UpLoadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	var cnt int = 0
	for _, file := range files {
		log.Println(file.Filename)
		dst := "./files/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			log.Println(err)
		} else {
			cnt++
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files are uploaded...", cnt))
}
