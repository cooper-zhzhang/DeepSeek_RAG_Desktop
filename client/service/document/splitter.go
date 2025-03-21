package document

import "github.com/tmc/langchaingo/textsplitter"

func GetDefaultPDFSplitter() *textsplitter.RecursiveCharacter {
	splitter := textsplitter.NewRecursiveCharacter(textsplitter.WithChunkSize(1000),
		textsplitter.WithChunkOverlap(100), textsplitter.WithSeparators([]string{"\n\n", "\n", "。", "!"}))
	return &splitter
}

func GetDefaultTextSplitter() *textsplitter.RecursiveCharacter {
	splitter := textsplitter.NewRecursiveCharacter(textsplitter.WithChunkSize(1000),
		textsplitter.WithChunkOverlap(10), textsplitter.WithSeparators([]string{"\n\n", "\n", "。", "!"}))
	return &splitter
}
