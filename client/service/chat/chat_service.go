package chat

import (
	"bufio"
	"bytes"
	"context"
	"dp_client/global"
	"dp_client/storage"
	"dp_client/storage/model"
	"encoding/json"
	"fmt"

	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type ChatService struct {
	url            string
	debug          bool
	conversationId uint64
	// TODO: ReqParamBuilder ReqParamBuilder
}

func NewChatService(url string, conversationId uint64, debug bool) *ChatService {
	return &ChatService{
		url:            url,
		debug:          debug,
		conversationId: conversationId,
	}
}

type ChatMessage struct {
	Role      string
	Content   string
	ToolCalls string
}

type ChatChunkResp struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done  bool `json:"done"`
	Error string
}

func (receiver *ChatService) init(url string) {
	receiver.url = url
}

func (receiver *ChatService) CreateHttpReq(ctx context.Context, content string) *http.Request {
	req := NewReqParamBuilder(ctx, global.MODEL_NAME, receiver.conversationId).
		buildReq(content, "你扮演一个回答问题的机器人，使用很热情的语句回答问题,尽量使用中文")

	global.Slog.DebugContext(ctx, "req", slog.Any("req", req))

	jsonData, err := json.Marshal(req)
	if err != nil {
		global.Slog.ErrorContext(ctx, "json.Marshal failed", slog.Any("err", err))
		return nil
	}

	// 创建一个 POST 请求
	httpReq, err := http.NewRequest("POST", receiver.url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.Slog.ErrorContext(ctx, "Error creating request: failed", slog.Any("err", err))
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
		global.Slog.ErrorContext(ctx, "Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	err = receiver.insertMessage(ctx, global.RoleUser, content, receiver.conversationId)
	if err != nil {
		global.Slog.ErrorContext(ctx, "insertMessage failed", slog.Any("err", err))
		return err
	}

	receiver.readChatStreamBody(ctx, resp.Body)

	return nil
}

func parseChatChunk(ctx context.Context, line string) *ChatChunkResp {
	var chunk *ChatChunkResp
	err := json.Unmarshal([]byte(line), &chunk)
	if err != nil {
		global.Slog.ErrorContext(ctx, "Unmarshal ChatChunkResp failed", slog.Any("err", err))
		return nil
	}

	if len(chunk.Error) != 0 {
		global.Slog.ErrorContext(ctx, "Error parsing JSON:", slog.Any("err", chunk.Error))
		return nil
	}

	return chunk
}

func (receiver *ChatService) insertMessage(ctx context.Context, role global.RoleType, content string, conversationId uint64) error {
	message := model.Message{
		Content:        content,
		Role:           string(role),
		ConversationId: conversationId,
	}

	return storage.NewMessageStorage(ctx).CreateMessage(&message)
}

// TODO:将resp解析为think部分和answer部分
func delThinkContent(sb strings.Builder) string {
	// 定义正则表达式，匹配 <think>...</think> 标签对及其内容
	re := regexp.MustCompile(`<think>[\s\S]*?</think>`)
	// 使用正则表达式替换匹配的内容为空字符串
	result := re.ReplaceAllString(sb.String(), "")

	return result
}
func (receiver *ChatService) readChatStreamBody(ctx context.Context, body io.ReadCloser) strings.Builder {

	sb := strings.Builder{}
	defer func() {
		if sb.Len() > 0 {
			err := receiver.insertMessage(ctx, global.RoleAssistant, delThinkContent(sb), receiver.conversationId)
			if err != nil {
				global.Slog.ErrorContext(ctx, "insertMessage failed", slog.Any("err", err))
			}
		}
	}()

	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		chunk := parseChatChunk(ctx, line)
		if chunk == nil {
			continue
		}

		if receiver.debug {
			fmt.Printf("Received resp: %s\n", chunk.Message.Content)
		} else {
			fmt.Printf(chunk.Message.Content)
		}

		// TODO： 用户态 根据配置 过滤掉 think
		sb.WriteString(chunk.Message.Content)

		if chunk.Done {
			if receiver.debug {
				fmt.Println("Done")
			} else {
				fmt.Println()
			}
			global.Slog.DebugContext(ctx, "chat resp", slog.Any("last chunk", chunk))
		}
	}

	return sb
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
