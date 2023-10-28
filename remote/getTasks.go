package remote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var url = "http://localhost:8080"

func GetAllItems() []TodoItem {

	resp, err := http.Get(url + "/getAllItems")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var items []TodoItem
	if err != json.Unmarshal(body, &items) {
		panic(err)
	}

	return items
}

func GetAllItemsInGroup(group string) []TodoItem {

	actualUrl := url + "/getAllItemsInGroup"
	var jsonInput = []byte(`{"group":"` + group + `"}`)
	req, err := http.NewRequest("GET", actualUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var items []TodoItem
	if err != json.Unmarshal(body, &items) {
		panic(err)
	}

	return items
}

func MarkItemComplete(id string) {
	actualUrl := url + "/markItemComplete"
	var jsonInput = []byte(`{"_id":"` + id + `"}`)
	req, err := http.NewRequest("POST", actualUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	client.Do(req)

}

func AddItem(item TodoItem) {
	actualUrl := url + "/addItem"
	jsonInput, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", actualUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	client.Do(req)

}
