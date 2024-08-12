package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Create) Note(entityId string, entityType string, note models.Note) (out models.RequestResponse, err error) {
	c.api.log("CreateTag request started...")

	if entityId == "" {
		err = fiber.ErrBadRequest
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: noEntityURL + entityType + "/" + entityId + notesURL,
		In:      note,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
