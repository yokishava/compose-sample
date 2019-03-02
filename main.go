package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wakashiyo/compose-sample/users"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Wolrd")
	})

	v1 := router.Group("api/v1")
	{
		v1.POST("/user", users.Register)
	}

	router.Run(":9000")
}
