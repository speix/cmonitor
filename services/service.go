package services

/*
	Provider abstraction to decouple models from service implementation
	and to easily add more providers (ex. MailChimp) if needed.
*/

import (
	"net/http"
	"bytes"
)

type EmailService interface {
	GetLists() (*http.Response, error)
	GetSubscribersFromList(list string) (*http.Response, error)
	AddSubscriber(list string, payload *bytes.Buffer) (*http.Response, error)
	DeleteSubscriber(list, email string) (*http.Response, error)
}
