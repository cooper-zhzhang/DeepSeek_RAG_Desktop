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
	llmEmbedded, err := embeddings.NewEmbedder(ollama_agent.GetLLMClient(ollama_agent.NOMIC_EMBED_TEXT))
	if err != nil {
		panic(err)
	}
	return llmEmbedded
}
