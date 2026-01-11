package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"errors"
)


// Calls getLinksLocal to get the names of all the tools in tools/
func initLinks() []string {
	var links []string
	if l, err := getLinksLocal(); err != nil {
		links = []string{ "markdown-renderer", "json-formatter" }
	} else {
		links = l
	}
	return links
}

// Searches the tools directory and returns all subdirectory names
func getLinksLocal() ([]string, error) {
	dirnames := []string{}
	dirs, err := os.ReadDir("./tools/")
	if err != nil {
		fmt.Println("Error reading filesystem")
		return []string{}, errors.New("error reading filesystem")
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name())
		dirnames = append(dirnames, dir.Name())
	}
	return dirnames, nil
}

// returns the file located at ./tools/{name}/dist/index.html
func getPage(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("./tools/" + name + "/dist/index.html")
	}
}

// Get the main landing page (.html) that will link all the separate tools together
func getLandingPage(c *gin.Context) {
	c.File("./client/dist/index.html")
}

func getMarkdownPage(c *gin.Context) {
	c.File("./tools/markdown-renderer/dist/index.html")
}

func getAssets(c *gin.Context) {
	toolname := c.Param("tool")
	assetname := c.Param("assetname")
	c.File("./tools/" + toolname + "/dist/assets/" + assetname)
}

func getAllPastebin(c *gin.Context) {
	entries := []string{
		"Some text",
		"some more text",
		"alot more text",
	}
	c.JSON(200, entries)
}

func getLinks(c *gin.Context) {
	dirnames := []string{}
	dirs, err := os.ReadDir("./tools/")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error reading filesystem" })
		return
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name())
		dirnames = append(dirnames, dir.Name())
	}
	c.JSON(200, dirnames)
}
