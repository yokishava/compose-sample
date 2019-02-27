package users

import (
  "net/http"
  "github.com/gin-gonic/gin"

  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
)

type Users struct {
  ID    int    `xorm:"id"`
  Name  string `xorm:"name"`
  Email string `xorm:"email"`
}

func SignIn(c *gin.Context) {
  //c.JSON(http.StatusOK, gin.H{
  //  "status": "posted",
  //  "message": "hello world!!!!!!!!!!",
  //  "nick": "nicknicknick",
  //})
  engine, _ := xorm.NewEngine("mysql", "test:test@tcp(db:3306)/test")
  user := Users{Name: "yoshikawa", Email: "yoshikawa@mail.com"}
  _, err := engine.Insert(&user)
  if err != nil {
    c.JSON(http.StatusOK, gin.H{
      "status": "error",
      "message": err,
    })
  } else {
    c.JSON(http.StatusOK, gin.H{
      "status": "posted",
      "message": "success!!!!",
    })
  }
}
