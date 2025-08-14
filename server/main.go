package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)


// ???? organize routes by GET,POST,PUT,DELETE or put all routes for a specific tool together?????
func main() {

	router := gin.Default()
	
	// Static for main landing pagee
	router.Static("/assets", "../client/dist/assets/")

	//====== GET ROUTES ========
	router.GET("/", getLandingPage)
	router.GET("/markdownrender", getMarkdownPage)
	// router.Static("/tools/markdown-renderer/assets", "../tools/markdown-renderer/dist/assets")
	router.GET("/tools/:tool/assets/*assetname", getAssets)



	//====== POST ROUTES ========



	//====== PUT ROUTES ========



	//====== DELETE ROUTES ========
	router.Run()
}
