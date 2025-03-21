package ollama_agent

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

func GetLLMCallBackHandler() callbacks.Handler {
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

}
func (handler *DeepSeekCallBacksHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {

}
func (handler *DeepSeekCallBacksHandler) HandleLLMError(ctx context.Context, err error) {

}
func (handler *DeepSeekCallBacksHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {

}
func (handler *DeepSeekCallBacksHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {

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
