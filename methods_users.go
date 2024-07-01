package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Get) Users(id string, params *Params) (user []models.User, err error) {
	c.api.log("GetUsers request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: usersURL,
		In:      nil,
		Out:     &models.User{},
		Params:  params,
	}

	if id != "" {
		options.BaseURL += "/" + id
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		user = []models.User{*options.Out.(*models.User)}
		c.api.log("returning the struct...")
		return
	} else {
		options.Out = &models.RequestResponse{}
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		user = options.Out.(*models.RequestResponse).Embedded.Users
		c.api.log("returning the struct...")
		return
	}
}
