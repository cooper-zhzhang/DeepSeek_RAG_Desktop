package document

import (
	"context"
	"dp_client/global"
	"log/slog"
	"os"

	"github.com/tmc/langchaingo/vectorstores"

	"github.com/tmc/langchaingo/vectorstores/qdrant"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

type FileType int16

const NullFileType FileType = 0
const TextFileType FileType = 1
const PDFFileType FileType = 2

type FileService struct {
	fileType FileType
	filePath string
}

func NewFileService() *FileService {
	return &FileService{}
}

type TextFile struct {
}

type PDFFile struct {
}

func (receiver *FileService) TextToChunks(ctx context.Context, filePath string) ([]schema.Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		global.Slog.Error("Open file failed", slog.Any("err", err))
		return nil, err
	}

	docLoaded := documentloaders.NewText(file)
	split := textsplitter.NewRecursiveCharacter()
	split.ChunkSize = 10000
	split.ChunkOverlap = 100
	docs, err := docLoaded.LoadAndSplit(ctx, split)
	if err != nil {
		global.Slog.Error("Load and split file failed", slog.Any("err", err))
		return nil, err
	}
	return docs, nil
}

func (receiver *FileService) storeDocs(ctx context.Context, docs []schema.Document, store *qdrant.Store) error {
	if len(docs) > 0 {
		_, err := store.AddDocuments(ctx, docs)
		if err != nil {
			global.Slog.ErrorContext(ctx, "AddDocuments failed", slog.Any("err", err))
			return err
		}
	}

	return nil
}

func (receiver *FileService) useRetriever(ctx context.Context, store *qdrant.Store, prompt string, topK int) ([]schema.Document, error) {

	optionsVector := []vectorstores.Option{
		vectorstores.WithScoreThreshold(0.80),
	}

	retriever := vectorstores.ToRetriever(store, topK, optionsVector...)

	doRetriever, err := retriever.GetRelevantDocuments(ctx, prompt)
	if err != nil {
		global.Slog.ErrorContext(ctx, "GetRelevantDocuments failed", slog.Any("err", err))
		return nil, err
	}

	return doRetriever, nil
}

/*
func GetAnswer(ctx context.Context, llm llms.Model, docRetrieved []schema.Document, input string) (string, error) {

	// 创建一个新的聊天消息历史记录
	history := memory.NewChatMessageHistory()
	// 将检索到的文档添加到历史记录中
	for _, doc := range docRetrieved {
		history.AddAIMessage(ctx, doc.PageContent)
	}
	// 使用历史记录创建一个新的对话缓冲区
	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))

	executor := agents.NewExecutor(
		agents.NewConversationalAgent(llm, nil, agents.WithCallbacksHandler(ollama_agent.GetLLMCallBackHandler())),
		//agents.WithOutputKey("text"),
		agents.WithMemory(conversation),
	)

	// 设置链调用选项
	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
	}
	res, err := chains.Run(ctx, executor, input, options...)
	if err != nil {
		return "", err
	}

	return res, nil
}
*/
