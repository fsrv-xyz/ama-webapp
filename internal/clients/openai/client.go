package openai

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Complete(prompt string) (*Response, error) {
	model := Model{
		Model:       "text-davinci-003",
		Prompt:      prompt,
		Temperature: 0.5,
		MaxTokens:   2000,
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(model)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", buf)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, requestError := http.DefaultClient.Do(req)
	if requestError != nil {
		log.Println(requestError)
		return nil, requestError
	}

	var r Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &r, nil
}
