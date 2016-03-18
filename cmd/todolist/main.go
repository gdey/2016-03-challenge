package main

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"github.com/gdey/todolist/task"
)

const (
	port = 8080
	PathIdx = 2
)

func main() {

	http.HandleFunc("/task/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d"), nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split("/", r.URL.Path)
	switch paths[PathIdx] {
	case "add":
		add(w, r)
	case "list":
		list(w, r)
	case "delete":
		delete(w, r)
	default:
		id(w, r)
	}
	return
}

func add(w http.ResponseWriter, r *http.Request) {

}

func list(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func id(w http.ResponseWriter, r *http.Request) {

    id := ""

    switch r.Method {
        case "GET": {
            t, err := task.Get(id)
            
        }
        case "PUT": {

            b, err := ioutil.ReadAll(r.Body)
            // TODO(nmeverden) check err
            defer r.Body.Close()
            
            
            // func Unmarshal(data []byte, v interface{}) error
            var t T
            if err := json.Unmarshal(b, t); err != nil {
                // TODO(nmeverden) handle err    
            }
            
            t.Write()
        }
    }
    
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Unsupported HTTP Method!"))
}
