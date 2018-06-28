package users_managers

import (
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)
import (
	conn "../conn"
)

func GetUsers(c *gin.Context) {
	db, err := conn.CreateDatabaseConnection()
	if err != nil {
		panic(err.Error())
	}

	var users []conn.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func CreateUser(c *gin.Context) {
	db, err := conn.CreateDatabaseConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.Exec("SET search_path TO golang")
	// get body
	user := conn.User{}
	c.BindJSON(&user)

	//update db
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
}
