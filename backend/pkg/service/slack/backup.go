package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	OK       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

func main() {
	fmt.Println("Creating user...")

	apiUrl := "https://slack.com/api/conversations.list"

	// create new http request
	request, error := http.NewRequest("GET", apiUrl, nil)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer token")

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	var res Response
	err := json.Unmarshal(responseBody, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// clean up memory after execution
	defer response.Body.Close()
}
