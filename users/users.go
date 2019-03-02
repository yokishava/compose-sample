package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Users user info
type Users struct {
	ID    int    `xorm:"id"`
	Name  string `xorm:"name"`
	Email string `xorm:"email"`
}

//Register register user info to db
func Register(c *gin.Context) {

	engine, _ := xorm.NewEngine("mysql", "test:test@tcp(db:3306)/test")

	name := c.Query("name")
	mail := c.Query("mail")

	user := Users{Name: name, Email: mail}

	_, err := engine.Insert(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "success regiter user",
		})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"name":   user.Name,
	// 	"mail":   user.Email,
	// })
}
