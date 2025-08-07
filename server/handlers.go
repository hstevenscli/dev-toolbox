package main

import (
	"github.com/gin-gonic/gin"
)


// Get the main landing page (.html) that will link all the separate tools together
func getLandingPage(c *gin.Context) {
	c.JSON(200, gin.H{"status": "HELLO WORLD"})
}
