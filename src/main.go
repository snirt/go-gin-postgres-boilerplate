package main

import (
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)
import (
	consts "./consts"
	paths "./paths"
	users_manager "./managers"
	reports_manager "./reports_manager"
)

// MAIN FUNCTION
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			consts.MESSAGE: "Pong!",
		})
	})
	// user services
	r.GET(	paths.USER, func(c *gin.Context) { users_manager.GetUsers(c) })
	r.POST(	paths.USER, func(c *gin.Context) { users_manager.CreateUser(c) })

	// reports services
	r.GET( paths.EXCEL, func(c *gin.Context) { reports_manager.GetExcel(c) })

	r.Run(":5000") // listen and serve on 0.0.0.0:5000
}

