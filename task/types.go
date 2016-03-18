package task

import (
	"encoding/json"
	"io/ioutil"
	"log -
	"os"
	"time"
)

type T struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Due         time.Time `json:"due"`
	Completed   bool      `json:"completed"`
}

const Rootdir = "task"
const FileTemplate = "task/%v.json"

var InvalidError = errors.Errorf("Invalid Object")

// Retrieve the task from the disk.
func Get(id int) (*T, error) {
	filename := fmt.Sprintf("task/%d.json", id)
	//file, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return nil, err
	//}
	//var t T
	//if err := json.Unmarshal(file, &t); err != nil {
	//	return nil, err
	//}
	//return &t, nil

	t, ReadUnmarshalFile(filename) (*T, error)
}

func (t *T) filename() string {
	id := 0
	if t != nil {
		id = t.Id
	}
	return fmt.Sprtinf(FileTemplate, id)
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
	return ioutil.WriteFile(t.filename(),bytes,os.ModePerm)
}

// Delete will delete the file from the disk.
func (t *T) Delete() (bool, error) {
	if t == nil {
		return false, InvalidError
	}
	if err := os.Remove(t.filename()); err != nil {
		return false, err
	}
	return true, nil
}

func generateId() int {
	return int(time.Now().Unix())
}

func List() ([]T, error) {

	var results []T

	infos, err := ioutil.ReadDir(Rootdir)
	if err != nil {
		return results, err
	}

	for _, info := range infos {
		b, err := ioutil.ReadFile(info.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		var t T
		if err = json.Unmarshal(b, &t); err != nil {
			results = append(results, t)
		}
	}
	return results, nil
}


func ReadUnmarshalFile(string filename) (*T, error) {
	file_content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var t T
	if err = json.Unmarshal(file_content, &t); err != nil {
		return nil, err
	}
	return &t, nil
	
	//Get function variant:
	//file, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return nil, err
	//}
	//var t T
	//if err := json.Unmarshal(file, &t); err != nil {
	//	return nil, err
	//}
	//return &t, nil
}
