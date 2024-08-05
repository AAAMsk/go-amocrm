package amocrm

import (
	"github.com/AAAMsk/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"regexp"
	"strconv"
)

func (c *Get) GetContactsByCustomFields(phone string, email string, params *Params) (isFind bool, contact models.Contact, err error) {
	c.api.log("GetContactsByCustomFields request is started...")

	var out []models.Contact
	regex := regexp.MustCompile("^(\\+)|[^\\d\\n]")
	i := 1
	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: contactsURL,
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

		out = options.Out.(*models.RequestResponse).Embedded.Contacts
		for _, value := range out {
			for _, item := range value.CustomFieldsValues {
				if (item.FieldCode == "PHONE" && regex.ReplaceAllString(item.Values[0].Value, "") == phone) ||
					(item.FieldCode == "EMAIL" && item.Values[0].Value == email) {
					return true, value, nil
				}
			}
		}

		i++
	}
	return false, contact, nil
}

func (c *Get) Contacts(contactID string, params *Params) (out []models.Contact, err error) {
	c.api.log("GetContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: contactsURL,
		In:      nil,
		Out:     &models.Contact{},
		Params:  params,
	}

	if contactID != "" {
		options.BaseURL += "/" + contactID
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Contact{*options.Out.(*models.Contact)}
	}

	if contactID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Contacts
	}

	c.api.log("returning the struct...")
	return
}

func (c *Create) Contact(in []models.Contact) (out models.RequestResponse, err error) {
	c.api.log("CreateContact request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: contactsURL,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}

func (c *Update) Contacts(contactID string, in []models.Contact) (out []models.Contact, err error) {
	c.api.log("ModifyContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsURL,
		In:      in,
		Out:     &models.Contact{},
		Params:  nil,
	}

	if contactID != "" {
		options.BaseURL += "/" + contactID
		options.In = in[0]
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Contact{*options.Out.(*models.Contact)}
	}

	if contactID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Contacts
	}

	c.api.log("returning the struct...")
	return
}

func (c *Get) ContactChats(contactID, chatID string) (out models.RequestResponse, err error) {
	c.api.log("GetContactChats request is started...")

	p := &Params{
		ContactID: contactID,
		ChatID:    chatID,
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsChatURL,
		In:      nil,
		Out:     &out,
		Params:  p,
	}

	c.api.getAgent(options.Method, options.BaseURL, options.Params)
	//req.RequestURI()

	c.api.log("returning the struct...")
	return
}

func (c *Create) ConnectChatToContact(in *[]models.Chat) (out models.RequestResponse, err error) {
	c.api.log("ConnectChatToContact request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsURL,
		In:      in,
		Out:     &out,
		Params:  nil,
	}
	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
