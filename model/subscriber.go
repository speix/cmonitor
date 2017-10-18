package model

import (
	"encoding/json"
	"errors"
	"bytes"
	"github.com/speix/cmonitor/services"
)

type Subscriber struct {
	List List `json:"list"`
	EmailAddress string `json:"emailaddress"`
	Name string `json:"name"`
}

func (s Subscriber) Add(es services.EmailService) (error){

	b := new(bytes.Buffer)
	encoder := json.NewEncoder(b)
	err := encoder.Encode(s)
	if err != nil{
		return errors.New("Failed to encode request: " + err.Error())
	}

	_, err = es.AddSubscriber(s.List.ListID, b)
	if err != nil{
		return errors.New(err.Error())
	}

	return nil
}

func (s Subscriber) Delete(es services.EmailService) (error){

	_, err := es.DeleteSubscriber(s.List.ListID, s.EmailAddress)
	if err != nil{
		return errors.New(err.Error())
	}

	return nil
}
