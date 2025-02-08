package main

import (
	"net/http"
	_"strings"
	"os"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	
	router.Static("assets", "./assets")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "AG_all.html", nil)
	})

	smtp := os.Getenv("SMTP_HOST")
	smtpport := os.Getenv("SMTP_PORT")
	println(smtp + smtpport)
	
   port := os.Getenv("PORT")
   if port == "" {
	   port = "8080"
   }
   router.Run(":"+port)

} 