package app

import "github.com/vamsikrishnasiddu/text-processor/controllers"

func mapUrls() {
	// the post request will be passed to the controller
	router.POST("/data", controllers.ProcessText)

}
