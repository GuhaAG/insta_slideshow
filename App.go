package main

import (
	"github.com/GuhaAG/insta_slideshow/backend"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("frontend", true)))
	backend.V1(r.Group("/api/v1"))
	r.Run(":3000")
}
