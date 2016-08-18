package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID    int    "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

//curl -X POST -H "Content-Type: application/json" http://localhost:8080/api/post  -d "{\"first\":\"anjun\",\"email\":\"last\",\"name\":\"ff\"}"
func main() {
	url := "http://localhost:8080/api/post"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"username":"anjun","firt":"an","last":"jun","email":"a@g.com"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
