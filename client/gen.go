package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func readBody(body io.ReadCloser) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应体内容
	fmt.Println(string(content))
}

func parseGenChunk(line string) *GenChunkResponse {
	var chunk *GenChunkResponse
	err := json.Unmarshal([]byte(line), &chunk)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return chunk
}

func readStreamBody(body io.ReadCloser) {
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {

		line := scanner.Text()
		chunk := parseGenChunk(line)
		if chunk == nil {
			return
		}

		fmt.Printf("Received resp: %s\n", chunk.Response)

		if chunk.Done {
			fmt.Println("Done")
			return
		}
	}

}

type GenReq struct {
	Model  string
	Prompt string
	Stream bool
}
type GenChunkResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	DoneReason         string    `json:"done_reason"`
	Context            []int     `json:"context"`
	TotalDuration      int       `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int       `json:"eval_duration"`
}

func createGenReq() *http.Request {
	// 定义请求的参数
	reqParam := GenReq{
		Model:  MODEL_NAME,
		Prompt: "who are you ",
		Stream: true,
	}
	// 将参数序列化为 JSON
	jsonData, err := json.Marshal(reqParam)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil
	}

	// 创建一个 POST 请求
	req, err := http.NewRequest("POST", GEN_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// 设置请求头，指定内容类型为 JSON
	req.Header.Set("Content-Type", "application/json")
	return req
}

func sendReq(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	return resp, nil
}

func sendGenReq() {
	req := createGenReq()
	resp, err := sendReq(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
	} else {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
	}

	readStreamBody(resp.Body)
}
