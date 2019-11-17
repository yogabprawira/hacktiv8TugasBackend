package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func home(c *gin.Context) {

}

func contact(c *gin.Context) {

}

func contactPost(c *gin.Context) {

}

func login(c *gin.Context) {

}

func logout(c *gin.Context) {

}

func articles(c *gin.Context) {

}

func articlesId(c *gin.Context) {

}

func articlesIdPost(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/home", home)
	router.GET("/contact", contact)
	router.POST("/contact", contactPost)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.GET("/articles", articles)
	router.GET("/articles/:id", articlesId)
	router.POST("/articles/:id", articlesIdPost)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
