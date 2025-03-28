package document

import (
	"context"
	"dp_client/service/ollama_agent"
	"net/url"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

var GlobalQdrantStore *qdrant.Store

var QdrantMap map[string]*qdrant.Store

func init() {

	QdrantMap = make(map[string]*qdrant.Store)

	store, err := createQdrantStore(viper.GetString("qdrant.collection"))
	if err != nil {
		panic(err)
	}
	GlobalQdrantStore = store

	QdrantMap[viper.GetString("qdrant.collection")] = store
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

func GetQdrant(ctx context.Context, collectName string) (*qdrant.Store, error) {

	if collectName == "" {
		return GlobalQdrantStore, nil
	}

	vectorStore := QdrantMap[collectName]
	if vectorStore != nil {
		return vectorStore, nil
	}

	vectorStore, err := createQdrantStore(collectName)
	if err != nil {
		return nil, err
	}

	QdrantMap[collectName] = vectorStore
	return vectorStore, nil

}

func createQdrantStore(collectName string) (*qdrant.Store, error) {
	qdrantUrl := viper.GetString("qdrant.url")
	qdUrl, err := url.Parse(qdrantUrl)
	if err != nil {
		return nil, err
	}
	store, err := qdrant.New(
		qdrant.WithURL(*qdUrl),                 // 设置URL
		qdrant.WithAPIKey(""),                  // 设置API密钥
		qdrant.WithCollectionName(collectName), // 设置集合名称
		qdrant.WithEmbedder(LanguageToolEmbedded()),
	)
	if err != nil {
		return nil, err
	}

	QdrantMap[collectName] = &store
	return &store, nil
}
