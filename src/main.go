package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://jsonplaceholder.typicode.com"

func main() {

	urls := []string{
		"/todos/1",
		"/posts/1",
		"/comments/1",
		"/users/1",
	}

	for _, url := range urls {
		endPoint := baseURL + url
		fmt.Println(endPoint)
		GetJSON(endPoint)
	}
}

// GetJSON Gets the JSON strings
func GetJSON(url string) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
