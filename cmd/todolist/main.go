package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gdey/todolist/task"
)

const (
	port    = 8080
	PathIdx = 2
)

func main() {

	http.HandleFunc("/task/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

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
		ID(w, r)
	}
	return
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported HTTP Method!"))
		return
	}
	t, err := getTaskFromHttpBody(r)
	if err != nil {
		painc(err)
	}
	t.Id = task.GetNextId()
	t.Write()
	w.WriteHeader(http.StatusOK)
}

func list(w http.ResponseWriter, r *http.Request) {
	ids, _ := task.List()
	bytes, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func delete(w http.ResponseWriter, r *http.Request) {
	id := idFromPath(r.URL.Path)
	t := task.T{id: id}
	t.Delete()
	w.WriteHeader(http.StatusOK)
}

func idFromPath(path string) int {
	paths := strings.Split("/", r.URL.Path)
	id, _ := strconv.Atoi(paths[len(paths)-1]) // TODO(nmeverde) parse Id from path
	return id
}
func getTaskFromHttpBody(r *http.Request) (*T, err) {
	b, err := ioutil.ReadAll(r.Body)
	// TODO(nmeverden) check err
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var t T
	if err := json.Unmarshal(b, t); err != nil {
		return nil, err
	}
	return &t, nil
}

// id if for getting the object, or writing a new object.
func ID(w http.ResponseWriter, r *http.Request) {

	id := idFromPath(r.URL.Path)
	// tasks/:id
	switch r.Method {

	case "GET":
		t, err := task.Get(id)
		bytes, err := json.Marshal(t)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	case "PUT":
		t, err := getTaskFromHttpBody(r)
		if err != nil {
			// TODO(nmeverden) handle err
			panic(err)
		}
		t.Id = id
		t.Write()
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Unsupported HTTP Method!"))
}
