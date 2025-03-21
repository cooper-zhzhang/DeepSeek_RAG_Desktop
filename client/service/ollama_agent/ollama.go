package ollama_agent

import (
	"dp_client/global"

	"github.com/tmc/langchaingo/llms/ollama"
)

type LLMName string

const (
	DEEP_SEEK_MODEL_7       LLMName = "deepseek-r1:7b"
	DEEP_SEEK_MODEL_1_DOT_5 LLMName = "deepseek-r1:1.5b"
	MXBAI_EMBED_LARGE       LLMName = "mxbai-embed-large"
)

func GetLLMClient(modelName LLMName) *ollama.LLM {

	switch modelName {
	case DEEP_SEEK_MODEL_1_DOT_5:
		return deepSeek1DOT5BLLM
	case DEEP_SEEK_MODEL_7:
		return deepSeek7BLLM
	case MXBAI_EMBED_LARGE:
		return maxbaiEmbedLarge
	default:
		return nil
	}
}

var deepSeek1DOT5BLLM *ollama.LLM
var deepSeek7BLLM *ollama.LLM
var maxbaiEmbedLarge *ollama.LLM

func init() {
	var err error
	deepSeek1DOT5BLLM, err = ollama.New(
		ollama.WithModel(string(DEEP_SEEK_MODEL_1_DOT_5)),
		ollama.WithServerURL(global.OLLAMA_URL),
	)
	if err != nil {
		panic(err)
	}

	deepSeek7BLLM, err = ollama.New(
		ollama.WithModel(string(DEEP_SEEK_MODEL_7)),
		ollama.WithServerURL(global.OLLAMA_URL),
	)
	if err != nil {
		panic(err)
	}

	maxbaiEmbedLarge, err = ollama.New(
		ollama.WithModel(string(MXBAI_EMBED_LARGE)),
		ollama.WithServerURL(global.OLLAMA_URL),
	)
	if err != nil {
		panic(err)
	}
}
