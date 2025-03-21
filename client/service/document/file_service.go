package document

import (
	"context"
	"dp_client/global"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/tmc/langchaingo/vectorstores"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

//IFileService针对一个文件的处理
// DocumentService是一个dataset的概念
//DocumentService 一个文档的集合，可以理解为dataset， 包含多个文件

type FileType int16

const NullFileType FileType = 0
const TextFileType FileType = 1
const PDFFileType FileType = 2

type IFileService interface {
	TextToChunks(ctx context.Context) ([]schema.Document, error)
	OpenFile(ctx context.Context) error
	CloseFile(ctx context.Context) error
	SetSplitter(ctx context.Context, splitter *textsplitter.RecursiveCharacter)
	StoreDocs(ctx context.Context, docs []schema.Document) error
	UseRetriever(ctx context.Context, prompt string, topK int) ([]schema.Document, error)
}

type FileService struct {
	IFileService
	FileType   FileType
	FilePath   string
	FileHandle *os.File
	Splitter   *textsplitter.RecursiveCharacter
	SchemaDocs []schema.Document
}

func NewFileService(ctx context.Context, fileType FileType, filePath string) (IFileService, error) {
	var file IFileService
	switch fileType {
	case TextFileType:
		file = &TextFile{
			FileService: FileService{
				FilePath: filePath,
				FileType: fileType,
				Splitter: GetDefaultTextSplitter(),
			},
		}
	case PDFFileType:
		file = &PDFFile{
			FileService{
				FilePath: filePath,
				FileType: fileType,
				Splitter: GetDefaultPDFSplitter(),
			},
		}
	default:
		return nil, errors.New("file type err")
	}

	return file, file.OpenFile(ctx)
}

type TextFile struct {
	FileService
}

type PDFFile struct {
	FileService
}

func (receiver *FileService) OpenFile(ctx context.Context) (err error) {
	receiver.FileHandle, err = os.Open(receiver.FilePath)
	if err != nil {
		global.Slog.ErrorContext(ctx, "Open file failed", slog.Any("err", err))
		return err
	}

	return nil
}

func (receiver *FileService) CloseFile(ctx context.Context) error {
	if receiver.FileHandle != nil {
		return receiver.FileHandle.Close()
	}
	return nil
}

func (receiver *FileService) DoAllThing(ctx context.Context) (err error) {
	err = receiver.OpenFile(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err := receiver.CloseFile(ctx)
		if err != nil {
			global.Slog.ErrorContext(ctx, "Close file failed", slog.Any("err", err))
		}
	}()

	_, err = receiver.TextToChunks(ctx)
	if err != nil {
		global.Slog.ErrorContext(ctx, "TextToChunks failed", slog.Any("err", err))
		return err
	}
	err = receiver.StoreDocs(ctx, receiver.SchemaDocs)
	if err != nil {
		global.Slog.ErrorContext(ctx, "StoreDocs failed", slog.Any("err", err))
		return err
	}

	docs, err := receiver.UseRetriever(ctx, "小明干什么的", 10)
	if err != nil {
		global.Slog.ErrorContext(ctx, "useRetriever failed", slog.Any("err", err))
		return err
	}

	fmt.Println(docs)

	return nil
}

func (receiver *FileService) TextToChunks(ctx context.Context) ([]schema.Document, error) {
	return nil, errors.New("need imple")
}

func (receiver *PDFFile) TextToChunks(ctx context.Context) ([]schema.Document, error) {
	fileInfo, err := receiver.FileHandle.Stat()
	if err != nil {
		global.Slog.Error("Get file info failed", slog.Any("err", err))
		return nil, err
	}

	fileSize := fileInfo.Size()
	docLoaded := documentloaders.NewPDF(receiver.FileHandle, fileSize)
	split := receiver.getSplitter(ctx)
	docs, err := docLoaded.LoadAndSplit(ctx, split)
	if err != nil {
		global.Slog.Error("Load and split file failed", slog.Any("err", err))
		return nil, err
	}

	receiver.SchemaDocs = docs
	return docs, nil
}

func (receiver *TextFile) TextToChunks(ctx context.Context) ([]schema.Document, error) {
	docLoaded := documentloaders.NewText(receiver.FileHandle)
	split := receiver.getSplitter(ctx)
	docs, err := docLoaded.LoadAndSplit(ctx, split)
	if err != nil {
		global.Slog.Error("Load and split file failed", slog.Any("err", err))
		return nil, err
	}

	receiver.SchemaDocs = docs
	return docs, nil
}

func (receiver *FileService) StoreDocs(ctx context.Context, docs []schema.Document) error {
	if len(docs) <= 0 {
		docs = receiver.SchemaDocs
	}
	if len(docs) > 0 {
		_, err := GlobalQdrantStore.AddDocuments(ctx, docs)
		if err != nil {
			global.Slog.ErrorContext(ctx, "AddDocuments failed", slog.Any("err", err))
			return err
		}
	}

	return nil
}

func (receiver *FileService) UseRetriever(ctx context.Context, prompt string, topK int) ([]schema.Document, error) {

	optionsVector := []vectorstores.Option{
		vectorstores.WithScoreThreshold(0.80),
	}

	retriever := vectorstores.ToRetriever(GlobalQdrantStore, topK, optionsVector...)

	doRetriever, err := retriever.GetRelevantDocuments(ctx, prompt)
	if err != nil {
		global.Slog.ErrorContext(ctx, "GetRelevantDocuments failed", slog.Any("err", err))
		return nil, err
	}

	return doRetriever, nil
}
func (receiver *FileService) SetSplitter(ctx context.Context, splitter *textsplitter.RecursiveCharacter) {

	if splitter == nil {
		global.Slog.WarnContext(ctx, "splitter is nil")
	}
	receiver.Splitter = splitter

}

func (receiver *FileService) getSplitter(ctx context.Context) *textsplitter.RecursiveCharacter {
	if receiver.Splitter == nil {
		newSplitter := textsplitter.NewRecursiveCharacter()
		receiver.Splitter = &newSplitter
	}
	return receiver.Splitter
}
