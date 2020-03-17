package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/opengraph"
)

type ErrorJson struct {
	ErrorMessage string `json:"ErrorMessage"`
}

func ogpParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Query("url")
		og, err := opengraph.Fetch(url)
		if err != nil {
			log.Println(err)
			eJson := ErrorJson{ErrorMessage: fmt.Sprint(err)}
			ctx.JSON(http.StatusOK, eJson)
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
