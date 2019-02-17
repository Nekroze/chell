package repl

import (
	prompt "github.com/c-bata/go-prompt"

	"github.com/Nekroze/chell/pkg/parsing"
)

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(
		[]prompt.Suggest{},
		d.GetWordBeforeCursor(),
		true,
	)
}

func buildPrompt() *prompt.Prompt {
	return prompt.New(
		parsing.Parse,
		completer,
		prompt.OptionTitle("chell"),
	)
}

func Run() {
	buildPrompt().Run()
}
