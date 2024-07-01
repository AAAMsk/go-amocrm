package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (c *Get) PipelineStatuses(id string, pipelineId string, params *Params) (status []models.PipelineStatus, err error) {
	c.api.log("PipelineStatuses request is started...")

	pipelineStatusUrl := strings.Replace(pipelineStatusesURL, "{pipeline_id}", pipelineId, -1)

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: pipelineStatusUrl,
		In:      nil,
		Out:     &models.PipelineStatus{},
		Params:  params,
	}

	if id != "" {
		options.BaseURL += "/" + id
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		status = []models.PipelineStatus{*options.Out.(*models.PipelineStatus)}
		c.api.log("returning the struct...")
		return
	} else {
		options.Out = &models.RequestResponse{}
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		status = options.Out.(*models.RequestResponse).Embedded.PipelineStatus
		c.api.log("returning the struct...")
		return
	}
}
