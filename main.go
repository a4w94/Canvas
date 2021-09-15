package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")        //讀取html靜態資源
	server.Static("/assets", "./template/assets") //讀取css靜態資源
	server.GET("/login", LoginPage)
	server.Run(":8000")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}
