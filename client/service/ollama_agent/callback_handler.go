package ollama_agent

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

func GetDeepSeekLLMCallBackHandler() callbacks.Handler {
	return &DeepSeekCallBacksHandler{}
}

type DeepSeekCallBacksHandler struct {
	callbacks.Handler
}

func (handler *DeepSeekCallBacksHandler) HandleText(ctx context.Context, text string) {
	fmt.Println(text)
}
func (handler *DeepSeekCallBacksHandler) HandleLLMStart(ctx context.Context, prompts []string) {
	fmt.Println(prompts)
}
func (handler *DeepSeekCallBacksHandler) HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent) {
	fmt.Println(ms)
}
func (handler *DeepSeekCallBacksHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {
	fmt.Println(res)
}
func (handler *DeepSeekCallBacksHandler) HandleLLMError(ctx context.Context, err error) {

}
func (handler *DeepSeekCallBacksHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {

}
func (handler *DeepSeekCallBacksHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {

	for k, v := range outputs {
		if k != "text" {
			continue
		}
		if str, ok := v.(string); ok {
			outputs[k] = "AI:" + str
		}
	}

}
func (handler *DeepSeekCallBacksHandler) HandleChainError(ctx context.Context, err error) {

}
func (handler *DeepSeekCallBacksHandler) HandleToolStart(ctx context.Context, input string) {

}
func (handler *DeepSeekCallBacksHandler) HandleToolEnd(ctx context.Context, output string) {

}
func (handler *DeepSeekCallBacksHandler) HandleToolError(ctx context.Context, err error) {

}
func (handler *DeepSeekCallBacksHandler) HandleAgentAction(ctx context.Context, action schema.AgentAction) {

}
func (handler *DeepSeekCallBacksHandler) HandleAgentFinish(ctx context.Context, finish schema.AgentFinish) {

}
func (handler *DeepSeekCallBacksHandler) HandleRetrieverStart(ctx context.Context, query string) {

}
func (handler *DeepSeekCallBacksHandler) HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document) {

}
func (handler *DeepSeekCallBacksHandler) HandleStreamingFunc(ctx context.Context, chunk []byte) {
}

func GetDeepSeekLLMStreamCallBackHandler(readChunk func(ctx context.Context, chunk []byte) error) callbacks.Handler {
	return &DeepSeekStreamCallBacksHandler{
		ReadChunk: readChunk,
	}
}

type DeepSeekStreamCallBacksHandler struct {
	callbacks.Handler
	ReadChunk func(ctx context.Context, chunk []byte) error
}

func (handler *DeepSeekStreamCallBacksHandler) HandleText(ctx context.Context, text string) {
	fmt.Println(text)
}
func (handler *DeepSeekStreamCallBacksHandler) HandleLLMStart(ctx context.Context, prompts []string) {
	fmt.Println(prompts)
}
func (handler *DeepSeekStreamCallBacksHandler) HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent) {
	fmt.Println(ms)
}
func (handler *DeepSeekStreamCallBacksHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {
	fmt.Println(res)
}
func (handler *DeepSeekStreamCallBacksHandler) HandleLLMError(ctx context.Context, err error) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {

	//fmt.Println(outputs)
	for k, v := range outputs {
		if k != "text" {
			continue
		}
		if str, ok := v.(string); ok {
			outputs[k] = "AI:" + str
		}
	}

}
func (handler *DeepSeekStreamCallBacksHandler) HandleChainError(ctx context.Context, err error) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleToolStart(ctx context.Context, input string) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleToolEnd(ctx context.Context, output string) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleToolError(ctx context.Context, err error) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleAgentAction(ctx context.Context, action schema.AgentAction) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleAgentFinish(ctx context.Context, finish schema.AgentFinish) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleRetrieverStart(ctx context.Context, query string) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document) {

}
func (handler *DeepSeekStreamCallBacksHandler) HandleStreamingFunc(ctx context.Context, chunk []byte) {
	_ = handler.ReadChunk(ctx, chunk)
}
