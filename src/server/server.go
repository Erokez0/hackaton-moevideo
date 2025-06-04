package server

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/Erokez0/hackaton-moevideo/src/categorizers/skydns"
	"github.com/Erokez0/hackaton-moevideo/src/config"
)

func categoriesHandler(c *gin.Context) {
	reqUrl := c.Query("url")
	if reqUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is required",
		})
		return
	}

	parsedUrl, err := url.ParseRequestURI(reqUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL",
		})
		return
	}

	reqUrl = parsedUrl.String()
	if r, err := http.Get(reqUrl); err != nil || r.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is unreachable",
		})
		return
	}

	
	result := []int{}
	var confident bool
	confidentQuery := c.Query("confident")
	if confidentQuery == "false" {
		confident = false
	} else {
		confident = true
	}

	result = append(result, skydns.Categorize(reqUrl, confident)...)

	c.JSON(http.StatusOK, gin.H{
		"categories": result,
	})
}

func Run() {
	r := gin.New()

	r.GET("/categories", categoriesHandler)

	address := config.ServerConfig.Address + ":" + config.ServerConfig.Port
	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}
	log.Println("\x1b[32mServer started\x1b[0m")
}
