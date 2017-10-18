package handlers

/*
	Handlers for receiving and processing requests from the client and responding accordingly.
*/

import (
	"net/http"
	"github.com/speix/cmonitor/services"
	"encoding/json"
	"github.com/speix/cmonitor/model"
)

type GetListsHandler struct {}
type GetListSubscribersHandler struct {}
type AddSubscriberHandler struct {}
type DeleteSubscriberHandler struct {}

type AddSubscriberRequestContainer struct {
	List string `json:"list"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type ServiceResponse struct {
	Message string `json:"message"`
}

func (h GetListsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	response := ServiceResponse{}
	
	cmService := services.CampaignMonitor{}

	lists := model.Lists{}
	err := lists.Get(cmService)
	if err != nil {
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	data, err := json.Marshal(lists)
	if err != nil {
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (h GetListSubscribersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	response := ServiceResponse{}

	list_id, _ := r.URL.Query()["list"]

	cmService := services.CampaignMonitor{}

	list := model.List{}
	list.ListID = list_id[0]
	err := list.GetSubscribersFromList(cmService)
	if err != nil {
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (h AddSubscriberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	response := ServiceResponse{}

	payload := AddSubscriberRequestContainer{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	cmService := services.CampaignMonitor{}

	subscriber := model.Subscriber{
		EmailAddress: payload.Email,
		Name: payload.Name,
		List: model.List{ListID: payload.List},
	}

	err = subscriber.Add(cmService)
	if err != nil{
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	data, err := json.Marshal(subscriber)
	if err != nil{
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (h DeleteSubscriberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	response := ServiceResponse{}

	list_id, _ := r.URL.Query()["list"]
	email, _ := r.URL.Query()["email"]

	cmService := services.CampaignMonitor{}

	subscriber := model.Subscriber{
		EmailAddress: email[0],
		List: model.List{ListID: list_id[0]},
	}

	err := subscriber.Delete(cmService)
	if err != nil{
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	data, err := json.Marshal(subscriber)
	if err != nil{
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(responseJson)

		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

