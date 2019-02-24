package repl

import (
	"os"
	"path/filepath"
	"strings"

	prompt "github.com/c-bata/go-prompt"

	"github.com/Nekroze/chell/pkg/parsing"
)

func getBinariesInPath() (binaries []prompt.Suggest) {
	foundBinaries := map[string]bool{}

	visitor := func(path string, f os.FileInfo, err error) error {
		if f != nil && !f.Mode().IsDir() {
			foundBinaries[f.Name()] = true
		}
		return err
	}

	for _, part := range strings.Split(os.Getenv("PATH"), ":") {
		filepath.Walk(part, visitor)
	}

	for name, _ := range foundBinaries {
		binaries = append(binaries, prompt.Suggest{name, ""})
	}

	return binaries
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(
		getBinariesInPath(),
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
