package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	// Although dependency injections are alternate way to create and pass instances but
	// for this we will go for globaly shared instance possible
	NewStore()
}

func main() {
	router := setup()

	// start the server
	log.Printf("Server Started at %v", PORT)
	router.Run(PORT)
}

func setup() *gin.Engine {
	router := createRouter()

	handleRequest(router)

	return router
}

// creates the router
func createRouter() *gin.Engine {
	return gin.Default()
}

// register all the url wrt their services
func handleRequest(router *gin.Engine) {
	router.POST(CREATE_SHORT_URL, CreateShortURL)
	router.GET(FETCH_SHORT_URL, FetchURL)
}
