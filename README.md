# Liaobots

liaobots网页sdk

## Getting started

```golang
package main

import (
	"fmt"
	"liaibots/client/request"
	"liaibots/client/sdk"

	uuid "github.com/satori/go.uuid"
)

func main() {
	client, err := sdk.NewClient("xxxxxxxxxxxxx")
	if err != nil {
		panic(err)
	}

	info, err := client.UserInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	models, err := client.Models()
	if err != nil {
		panic(err)
	}
	fmt.Println(models)

	req := request.ChatReq{
		ConversationID: uuid.NewV4().String(),
		Model: request.Model{
			ID:         "gpt-3.5-turbo-16k",
			Name:       "GPT-3.5-16k",
			MaxLength:  48000,
			TokenLimit: 16000,
		},
		Messages: []request.Message{
			{
				Role:    "user",
				Content: "你好",
			},
		},
		Prompt: "友好的问答机器人",
	}

	resp, err := client.Chat(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	messages := []request.Message{
		{
			Role:    "user",
			Content: "你好",
		},
		{
			Role:    "assistant",
			Content: resp,
		},
	}
	err = client.Recommend(messages)
	if err != nil {
		panic(err)
	}
}

```