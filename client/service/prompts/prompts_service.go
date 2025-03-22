package prompts

const DefaultPromptSuffix = `
	Begin!
	ChatByChains History:
	{{.history}}
	Follow Up Input: {{.input}}
	Standalone question:`
