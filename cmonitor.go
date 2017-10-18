package main

import (
	"fmt"
	"net/http"
	"github.com/speix/cmonitor/handlers"
)

 /*
	Basic function chaining for logging the requested handler in console.
 */
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Handler called: %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main(){

	lists := handlers.GetListsHandler{}
	listSubscribers := handlers.GetListSubscribersHandler{}
	addSubscriber := handlers.AddSubscriberHandler{}
	deleteSubscriber := handlers.DeleteSubscriberHandler{}

	server := &http.Server{
		Addr: ":8080",
	}

	http.Handle("/lists", log(lists))
	http.Handle("/subscribers", log(listSubscribers))
	http.Handle("/subscriber/add", log(addSubscriber))
	http.Handle("/subscriber/delete", log(deleteSubscriber))

	http.Handle("/", log(http.FileServer(http.Dir("./static"))))

	server.ListenAndServe()
}
