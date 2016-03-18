package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

const rootdir = "task"
const FileTemplate = rootdir + "/%v.json"

var InvalidError = errors.Errorf("Invalid Object")

// readFile will read the file from the filesystem, and return a T object if it exists.
func readFile(fn string) (*T, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var t T
	if err = json.Unmarshal(content, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func filenameForID(id int) { return fmt.Sprintf(FileTemplate, id) }

// Retrieve the task from the disk.
func Get(id int) (t *T, err error) {
	t, err = readFile(filenameForID(id))
	return
}

// GetNextId will get an id that is not being used.
func GetNextId() int {
	ids, _ := List()
	max := 0
	for _, i := range ids {
		if i > max {
			max = i
		}
	}
	return max + 1
}

// This will return a list of id's in the data store.
func List() ([]int, error) {

	var results []int

	names, _ := filepath.Glob(rootdir + "/*")
	for _, name := range names {
		dir, file := path.Split(name)
		ext := path.Ext(file)
		frune := []rune(file)
		f := frune[:len(frune)-len(ext)]
		id, err := strconv.Atoi(f)
		if err != nil {
			// Skip invalid id's on the disk
			log.Println("Invalid file id ", file)
			continue
		}
		results = append(results, id)
	}
	return results, nil
}

type T struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Due         time.Time `json:"due"`
	Completed   bool      `json:"completed"`
}

// Write will write the Task to disk.
func (t *T) Write() error {
	if t == nil {
		return InvalidError
	}
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(t.filename(), bytes, os.ModePerm)
}

// Delete will delete the file from the datastore.
func (t *T) Delete() (bool, error) {
	if t == nil {
		return false, InvalidError
	}
	if err := os.Remove(t.filename()); err != nil {
		return false, err
	}
	return true, nil
}

func (t *T) filename() string {
	if t != nil {
		return filenameForID(0)
	}
	return filenameForID(t.Id)
}
