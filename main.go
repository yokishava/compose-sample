package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/wakashiyo/compose-sample/users"
)

func main() {
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello Wolrd")
  })

  v1 := router.Group("api/v1")
  {
    v1.GET("/user", users.SignIn)
  }

  router.Run(":9000")
}
