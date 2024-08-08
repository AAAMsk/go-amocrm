package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

func (c *Get) GetCompaniesByCustomFields(queryParams []string) (isFind bool, company models.Company, err error) {
	c.api.log("GetCompaniesByCustomFields request is started...")

	var out []models.Company
	var p = Params{}

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: companiesURL,
		In:      nil,
		Out:     &models.RequestResponse{},
		Params:  &p,
	}

	regex := regexp.MustCompile("^(\\+)|[^\\d\\n]")
	isFind = false
	i := 0
	for _, param := range queryParams {
		options.Params.Query = param
		if err = c.api.makeRequest(options); err != nil {
			if err.Error() == "No content" {
				continue
			}
			return false, company, err
		}

		out = options.Out.(*models.RequestResponse).Embedded.Companies
		if len(out) > 1 {
			for _, value := range out {
				for _, item := range value.CustomFieldsValues {
					if i == 0 {
						if item.FieldCode == "EMAIL" {
							if len(queryParams) > 1 && item.Values[0].Value == queryParams[1] {
								return true, value, nil
							} else {
								return true, out[0], nil
							}
						} else {
							continue
						}
					} else if i == 1 {
						if item.FieldCode == "PHONE" {
							if regex.ReplaceAllString(item.Values[0].Value.(string), "") == queryParams[0] {
								return true, value, nil
							} else {
								return true, out[0], nil
							}
						} else {
							continue
						}
					}
				}
			}
		} else {
			return true, out[0], nil
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
