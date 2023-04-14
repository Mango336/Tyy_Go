package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloWeb() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("===Hello Web middleware...===")
	}
}
