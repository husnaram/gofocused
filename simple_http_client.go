package gofocused

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseURL = "http://localhost:9090"

type Student struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Grade int    `json:"grade,omitempty"`
}

func RunFetching() {
	var users, err = FetchUser()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	// fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	for _, each := range users {
		fmt.Printf("ID: \t%s\nName: \t%s\nGrade: \t%d\n-----\n", each.ID, each.Name, each.Grade)
	}
}

func FetchUser() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var data []Student

	request, err := http.NewRequest("POST", BaseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
