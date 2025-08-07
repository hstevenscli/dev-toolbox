package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)


// ???? organize routes by GET,POST,PUT,DELETE or put all routes for a specific tool together?????
func main() {

	router := gin.Default()

	//====== GET ROUTES ========
	router.GET("/", getLandingPage)



	//====== POST ROUTES ========



	//====== PUT ROUTES ========



	//====== DELETE ROUTES ========
	router.Run()
}
