package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"regexp"
	"strconv"
)

func (c *Get) GetCompaniesByCustomFields(phone string, email string, params *Params) (isFind bool, company models.Company, err error) {
	c.api.log("GetCompaniesByCustomFields request is started...")

	var out []models.Company
	regex := regexp.MustCompile("^(\\+)|[^\\d\\n]")
	i := 1
	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: companiesURL,
		In:      nil,
		Out:     &models.RequestResponse{},
		Params:  params,
	}

	for {
		options.Params.Page = strconv.Itoa(i)

		if err = c.api.makeRequest(options); err != nil {
			if err.Error() == "No content" {
				break
			}
			log.Println(err)
			return
		}

		out = options.Out.(*models.RequestResponse).Embedded.Companies
		for _, value := range out {
			for _, item := range value.CustomFieldsValues {
				if (item.FieldCode == "PHONE" && regex.ReplaceAllString(item.Values[0].Value.(string), "") == phone) ||
					(item.FieldCode == "EMAIL" && item.Values[0].Value == email) {
					return true, value, nil
				}
			}
		}

		i++
	}

	return false, company, nil
}

func (c *Get) Companies(companyID string, params *Params) (out []models.Company, err error) {
	c.api.log("Get Companies request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: companiesURL,
		In:      nil,
		Out:     &models.Company{},
		Params:  params,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	c.api.log("returning the struct...")
	return
}

func (c *Create) Companies(in []models.Company) (out []models.Company, err error) {
	c.api.log("Create Company request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	options.Out = &models.RequestResponse{}
	if err = c.api.makeRequest(options); err != nil {
		return
	}
	out = options.Out.(*models.RequestResponse).Embedded.Companies

	c.api.log("returning the struct...")
	return
}

func (c *Update) Companies(companyID string, in []models.Company) (out []models.Company, err error) {
	c.api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		options.In = in[0]

		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	c.api.log("returning the struct...")
	return
}
