package services

/*
	Campaign Monitor basic Lists/Subscribers implementation.
	Every functionality returns a direct response from the service along with an error struct if applicable.

	Set CM_API_CLIENT and CM_API_KEY as environment variables with values from the Campaign Monitor API.
*/

import (
	"net/http"
	"encoding/base64"
	"errors"
	"bytes"
	"encoding/json"
	"os"
)

type CampaignMonitor struct {}

type responseBody struct {
	Message string
}

func (cm CampaignMonitor) GetLists() (*http.Response, error){

	client := &http.Client{}
	url := "https://api.createsend.com/api/v3.1/clients/" + os.Getenv("CM_API_CLIENT") + "/lists.json"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("CM_API_KEY"))))

	if err != nil{
		return nil, errors.New("Failed to prepare request: " + err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Failed to execute request: " + err.Error())
	}

	if response.StatusCode != 200 {

		body := responseBody{}
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			return nil, errors.New(body.Message)
		}

		return nil, errors.New("Unable to Get Lists: " + response.Status)
	}

	return response, nil
}

func (cm CampaignMonitor) GetSubscribersFromList(list string) (*http.Response, error){

	client := &http.Client{}
	url := "https://api.createsend.com/api/v3.1/lists/" + list + "/active.json"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("CM_API_KEY"))))
	if err != nil{
		return nil, errors.New("Failed to prepare request: " + err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Failed to execute request: " + err.Error())
	}

	if response.StatusCode != 200 {

		body := responseBody{}
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			return nil, errors.New(body.Message)
		}

		return nil, errors.New("Unable to Get Lists: " + response.Status)
	}

	return response, nil
}

func (cm CampaignMonitor) AddSubscriber(list string, payload *bytes.Buffer) (*http.Response, error) {

	client := &http.Client{}
	url := "https://api.createsend.com/api/v3.1/subscribers/" + list + ".json"

	request, err := http.NewRequest("POST", url, payload)
	request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("CM_API_KEY"))))
	if err != nil{
		return nil, errors.New("Failed to prepare request: " + err.Error())
	}

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil{
		return nil, errors.New("Failed to handle request: " + err.Error())
	}

	if response.StatusCode != 201 {

		body := responseBody{}
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil{
			return nil, errors.New("Unable to Add Subscriber: " + response.Status)
		}

		return nil, errors.New(body.Message)
	}

	return response, nil
}

func (cm CampaignMonitor) DeleteSubscriber(list, email string) (*http.Response, error) {

	client := &http.Client{}
	url := "https://api.createsend.com/api/v3.1/subscribers/" + list +".json?email=" + email

	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("CM_API_KEY"))))
	if err != nil{
		return nil, errors.New("Failed to prepare request: " + err.Error())
	}


	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, errors.New("Failed to make request: " + err.Error())
	}

	if response.StatusCode != 200 {

		body := responseBody{}
		err = json.NewDecoder(response.Body).Decode(&body)
		if err != nil{
			return nil, errors.New("Unable to Delete Subscriber: " + response.Status)
		}

		return nil, errors.New(body.Message)
	}

	return response, nil
}