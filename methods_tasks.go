package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Create) Task(task []models.Task) (out models.RequestResponse, err error) {
	c.api.log("CreateTag request started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: noEntityURL + tasksURL,
		In:      task,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
