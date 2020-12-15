package gofocused

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

var jsonString = `{"Name": "Traumen Rings", "Age": 3200}`
var jsonData = []byte(jsonString)

func DecodeJSONtoVarObjStruct() {
	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User\t:", data.FullName)
	fmt.Println("Age\t:", data.Age)
}

// DecodeJSONtoInterface decoding JSON data to map[string]interface{} & interface{}
func DecodeJSONtoInterface() {
	var dataOne map[string]interface{}
	json.Unmarshal(jsonData, &dataOne)

	fmt.Println("User\t:", dataOne["Name"])
	fmt.Println("Age\t:", dataOne["Age"])

	// ---

	var dataTwo interface{}
	json.Unmarshal(jsonData, &dataTwo)
	var decodeData = dataTwo.(map[string]interface{})
	fmt.Println("User\t:", decodeData["Name"])
	fmt.Println("Age\t:", decodeData["Age"])
}

func JSONArrayDecodeToArrayObj() {
	var jsonArr = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	var data []User

	var err = json.Unmarshal([]byte(jsonArr), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User 1\t:", data[0].FullName)
	fmt.Println("User 2\t:", data[1].FullName)
}

func EncodeObjToJSONString() {
	var obj = []User{
		{"Hasan Basari", 39},
		{"Longman Hart", 26},
	}

	var jsonData, err = json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
