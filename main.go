package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/opengraph"
)

func ogpParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Query("url")
		og, err := opengraph.Fetch(url)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, nil)
		} else {
			ctx.JSON(http.StatusOK, og)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/ogp", ogpParser())
	router.Run(":30002")
}
