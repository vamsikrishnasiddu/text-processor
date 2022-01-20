package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vamsikrishnasiddu/text-processor/rest_errors"
	"github.com/vamsikrishnasiddu/text-processor/services"
)

// processText processes the text and returns the
// top 10 most used words with their occurence as a json key value pairs
func ProcessText(c *gin.Context) {
	// reading the text from the body
	rb, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		err := rest_errors.NewBadRequestError("invalid request body")
		c.JSON(err.Status(), err)
		return
	}

	var restErr rest_errors.RestErr

	text := string(rb)
	// passing the text to wordservice where actual business logic is implemented
	data, restErr := services.WordService.ProcessWordOccurence(text)

	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusCreated, data)

}
