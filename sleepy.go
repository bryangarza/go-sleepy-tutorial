package main

import (
	"net/http"
	"net/url"
)

type Resource interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

type (
	GetNotSupported    struct{}
	PostNotSupported   struct{}
	PutNotSupported    struct{}
	DeleteNotSupported struct{}
)

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

func (PostNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

func (PutNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

func (DeleteNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world."))
}

type API struct{}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		method := request.Method // Get HTTP method (string)
		request.ParseForm()      // Populates request.Form
		values := request.Form

		switch method {
		case "GET":
			code, data = resource.GET(values)
		case "POST":
			code, data = resource.Post(values)
		case "PUT":
			code, data = resource.Put(values)
		case "DELETE":
			code, data = resource.Delete(values)
		}

		// TODO: write response
	}
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":3000", nil)
}
