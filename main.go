package main

import (
	// "fmt"

	"os"
	"github.com/gin-gonic/gin"
)


// ???? organize routes by GET,POST,PUT,DELETE or put all routes for a specific tool together?????
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	
	// Static for main landing pagee
	router.Static("/assets", "./client/dist/assets/")

	//////// GET ROUTES ////////////// GET ROUTES ////////////// GET ROUTES ////////////// GET ROUTES ////////////// GET ROUTES ////////
	router.GET("/db/pastebin", getAllPastebin)

	router.GET("/", getLandingPage)

	// Autmoatically makes GET routes for each tool
	// get page will try to return the following file: ./tools/{toolname}/dist/index.html
	var links = initLinks()
	for i := range links {
		router.GET(links[i], getPage(links[i]))
	}
	router.GET("/links", getLinks)

	// router.Static("/tools/markdown-renderer/assets", "../tools/markdown-renderer/dist/assets")
	router.GET("/tools/:tool/assets/*assetname", getAssets)



	//////// POST ROUTES ////////////// POST ROUTES ////////////// POST ROUTES ////////////// POST ROUTES ////////////// POST ROUTES ////////



	//////// PUT ROUTES ////////////// PUT ROUTES ////////////// PUT ROUTES ////////////// PUT ROUTES ////////////// PUT ROUTES ////////////// PUT ROUTES ////////



	//////// DELETE ROUTES ////////////// DELETE ROUTES ////////////// DELETE ROUTES ////////////// DELETE ROUTES ////////////// DELETE ROUTES ////////
    router.Run("0.0.0.0:"+port)
}
