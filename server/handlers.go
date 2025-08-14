package main

import (
	"github.com/gin-gonic/gin"
)


// Get the main landing page (.html) that will link all the separate tools together
func getLandingPage(c *gin.Context) {
	c.File("../client/dist/index.html")
}

func getMarkdownPage(c *gin.Context) {
	c.File("../tools/markdown-renderer/dist/index.html")
}

func getAssets(c *gin.Context) {
	toolname := c.Param("tool")
	assetname := c.Param("assetname")
	c.File("../tools/" + toolname + "/dist/assets/" + assetname)
}
