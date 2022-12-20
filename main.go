package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/bonsai-oss/jsonstatus"
	"github.com/bonsai-oss/webbase"

	"ama-webapp/internal/clients/openai"
)

type Input struct {
	Text string `json:"text"`
}

//go:embed ui
var uiStorage embed.FS

func OpenAIHandlerBuilder(token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Input
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			jsonstatus.Status{Code: http.StatusBadRequest, Message: fmt.Sprintf("invalid input %+q", err)}.Encode(w)
			return
		}

		client := openai.NewClient(token)
		response, err := client.Complete(input.Text)
		if err != nil {
			jsonstatus.Status{Code: http.StatusInternalServerError, Message: fmt.Sprintf("invalid input %+q", err)}.Encode(w)
			return
		}

		jsonstatus.Status{Code: http.StatusOK, Message: fmt.Sprintf("%s", response.Choices[0].Text)}.Encode(w)
	}
}

func main() {
	router := webbase.NewRouter()
	token := os.Getenv("OPENAI_TOKEN")
	if token == "" {
		panic("OPENAI_TOKEN is not set")
	}

	sub, err := fs.Sub(uiStorage, "ui")
	if err != nil {
		panic(err)
	}
	router.Path("/api").Methods(http.MethodPost).HandlerFunc(OpenAIHandlerBuilder(token))
	router.PathPrefix("/").Methods(http.MethodGet).Handler(http.FileServer(http.FS(sub)))

	webbase.ServeRouter("ama", router, webbase.WithoutServiceEndpoint())
}
