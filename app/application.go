package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// urls will be mapped to the respective handlers and we start the http server on port 8080
func StartApplication() {
	mapUrls()
	//server is listening on port 8080
	router.Run(":8080")
}
