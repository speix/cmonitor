package model

import (
	"encoding/json"
	"errors"
	"github.com/speix/cmonitor/services"
)

type List struct {
	ListID string `json:"listid"`
	Name string `json:"name"`
	Subscribers []Subscriber `json:"subscribers"`
}

type Lists []List

type Container struct {
	Results []Subscriber `json:"results"`
}

func (lists *Lists) Get(es services.EmailService) (error){

	response, err := es.GetLists()
	defer response.Body.Close()
	if err != nil{
		return errors.New("Failed to request: " + err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&lists)
	if err != nil{
		return errors.New("Failed to decode response: " + err.Error())
	}

	return nil
}

func (list *List) GetSubscribersFromList(es services.EmailService) (error){

	response, err := es.GetSubscribersFromList(list.ListID)
	defer response.Body.Close()
	if err != nil{
		return errors.New("Failed to request: " + err.Error())
	}

	container := &Container{}
	err = json.NewDecoder(response.Body).Decode(&container)
	if err != nil{
		return errors.New("Failed to decode response: " + err.Error())
	}

	list.Subscribers = container.Results

	return nil
}
