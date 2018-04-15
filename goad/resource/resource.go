package resource

import (
	"encoding/xml"
	"net/http"
)

// Resource ...
type Resource interface {
	Get(r *http.Request) (Status, interface{})
	Post(r *http.Request) (Status, interface{})
}

// Base ...
type Base struct {
}

// Get ...
func (b Base) Get(r *http.Request) (Status, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

// Post ...
func (b Base) Post(r *http.Request) (Status, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

// XMLHandler ...
func XMLHandler(resource Resource) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var status Status
		var data interface{}
		switch r.Method {
		case "GET":
			status, data = resource.Get(r)
		case "POST":
			status, data = resource.Get(r)
		default:
			// TODO:
		}

		// Return Response
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Vary", "Accept-Encoding,User-Agent")
		var content []byte
		var e error
		if status.Success {
			content, e = xml.Marshal(data)
			w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		} else {
			content = []byte(status.Message)
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		}
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status.Code)
		w.Write(content)
	}
}
