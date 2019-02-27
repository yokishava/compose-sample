package users

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "status": "posted",
    "message": "hello world!!!!!!!!!!",
    "nick": "nicknicknick",
  })
}
