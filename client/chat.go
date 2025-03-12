package main

import (
	"bufio"
	"bytes"
	"context"
	"dp_client/global"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type ChatService struct {
	url   string
	debug bool
}

func NewChatService(url string, debug bool) *ChatService {
	return &ChatService{
		url:   url,
		debug: debug,
	}
}

type RoleType string

const (
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
)

type ChatReq struct {
	Model    string
	Messages []ChatMessage
	Stream   bool
}
type ChatMessage struct {
	Role      string
	Content   string
	ToolCalls string
}

func (receiver *ChatReq) buildReq(question string) *ChatReq {
	return &ChatReq{
		Model: MODEL_NAME,
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: "why is the sky blue?",
			},
			{
				Role:    "assistant",
				Content: "due to rayleigh scattering.",
			},
			{
				Role:    "user",
				Content: "how is that different than mie scattering?",
			},
			//
			// {
			// 	Role:      string(RoleUser),
			// 	Content:   "Do you know my name?",
			// 	ToolCalls: "",
			// },
			// {
			// 	Role:      string(RoleAssistant),
			// 	Content:   "I don’t have access to your personal information, including your name. However, if you need assistance with anything related to your name or other details, feel free to provide more context, and I’ll do my best to help!",
			// 	ToolCalls: "",
			// },
			// {
			// 	Role:      string(RoleUser),
			// 	Content:   "My name is Jerry.",
			// 	ToolCalls: "",
			// },
			// {
			// 	Role:      string(RoleAssistant),
			// 	Content:   "It seems like you're referring to someone named Jerry. If this is a conversation or a task where I can help, please feel free to provide more details, and I'll do my best to assist!",
			// 	ToolCalls: "",
			// },
			// {
			// 	Role:      string(RoleUser),
			// 	Content:   question,
			// 	ToolCalls: "",
			// }
		},
		Stream: true,
	}
}

type ChatChunkResp struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done bool `json:"done"`
}

func (receiver *ChatService) init(url string) {
	receiver.url = url
}

func (receiver *ChatService) CreateHttpReq(ctx context.Context, content string) *http.Request {
	chatReq := ChatReq{}
	req := chatReq.buildReq(content)

	jsonData, err := json.Marshal(req)
	if err != nil {
		global.SlogLogger.ErrorContext(ctx, "json.Marshal failed", slog.Any("err", err))
		return nil
	}

	global.SlogLogger.InfoContext(ctx, "req", slog.Any("req", req))

	// 创建一个 POST 请求
	httpReq, err := http.NewRequest("POST", receiver.url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.SlogLogger.ErrorContext(ctx, "Error creating request: failed", slog.Any("err", err))
		return nil
	}

	// 设置请求头，指定内容类型为 JSON
	httpReq.Header.Set("Content-Type", "application/json")
	return httpReq
}

func (receiver *ChatService) Run(ctx context.Context, content string) error {
	req := receiver.CreateHttpReq(ctx, content)
	resp, err := receiver.sendReq(req)
	if err != nil {
		global.SlogLogger.ErrorContext(ctx, "Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// TODO: go func write channel
	// TODO:将对话内容写入db
	receiver.readChatStreamBody(resp.Body)
	return nil
}

func parseChatChunk(line string) *ChatChunkResp {
	var chunk *ChatChunkResp
	err := json.Unmarshal([]byte(line), &chunk)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return chunk
}

func (receiver *ChatService) readChatStreamBody(body io.ReadCloser) {
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		chunk := parseChatChunk(line)
		if chunk == nil {
			return
		}

		if receiver.debug {
			fmt.Printf("Received resp: %s\n", chunk.Message.Content)
		} else {
			fmt.Printf(chunk.Message.Content)
		}

		if chunk.Done {
			if receiver.debug {
				fmt.Println("Done")
			}
			return
		}
	}
}

func (receiver *ChatService) sendReq(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err := resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return resp, nil
}
