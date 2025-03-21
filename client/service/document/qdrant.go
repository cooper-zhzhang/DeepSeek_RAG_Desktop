package document

import (
	"dp_client/service/ollama_agent"
	"net/url"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

var GlobalQdrantStore *qdrant.Store

func init() {

	// TODO: 多个集合
	qdrantUrl := viper.GetString("qdrant.url")
	qdUrl, err := url.Parse(qdrantUrl)
	if err != nil {
		panic(err)
	}
	store, err := qdrant.New(
		qdrant.WithURL(*qdUrl), // 设置URL
		qdrant.WithAPIKey(""),  // 设置API密钥
		qdrant.WithCollectionName(viper.GetString("qdrant.collection")), // 设置集合名称
		qdrant.WithEmbedder(LanguageToolEmbedded()),
	)
	if err != nil {
		panic(err)
	}
	GlobalQdrantStore = &store
}

func LanguageToolEmbedded() *embeddings.EmbedderImpl {
	//	MXBAI_EMBED_LARGE collections size=1024
	// NOMIC_EMBED_TEXT collections size=768
	llmEmbedded, err := embeddings.NewEmbedder(ollama_agent.GetLLMClient(ollama_agent.MXBAI_EMBED_LARGE))
	if err != nil {
		panic(err)
	}
	return llmEmbedded
}
