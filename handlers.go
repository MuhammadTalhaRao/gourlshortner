package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// it creates and stores shortcodes for urls
// NOTE: I have not checked if URL already exists becasue we have a constraint to create the same code each time
func CreateShortURL(c *gin.Context) {
	// get the url from req
	var request CreateShotURLRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, createErrResp(err.Error()))
	}

	log.Printf("[CreateShortURL] Recived request: %v", request)

	// create shortcode
	shortCode := createShortCode(request.URL)

	// store it
	Store.AddURLToStore(request.URL, shortCode)

	resp := createSuccessResp(shortCode)

	log.Printf("[CreateShortURL] Requtned response: %v", resp)

	// return resp
	c.JSON(http.StatusOK, resp)
}

// it returns the urls based on request short code after ensuring the availibility
func FetchURL(c *gin.Context) {
	// get the short code from req
	shortCode := c.Param("shortcode")

	log.Printf("[FetchURL] Recived request: %v", shortCode)

	// fetch the url from store
	url, err := Store.FetchURLFromStore(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	log.Printf("[FetchURL] Requtned response: %v", url)

	// return resp
	c.Redirect(http.StatusMovedPermanently, url)
}
