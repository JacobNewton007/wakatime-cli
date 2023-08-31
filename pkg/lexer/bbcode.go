package lexer

import (
	"github.com/wakatime/wakatime-cli/pkg/heartbeat"

	"github.com/alecthomas/chroma/v2"
)

// BBCode lexer.
type BBCode struct{}

// Lexer returns the lexer.
func (l BBCode) Lexer() chroma.Lexer {
	return chroma.MustNewLexer(
		&chroma.Config{
			Name:      l.Name(),
			Aliases:   []string{"bbcode"},
			MimeTypes: []string{"text/x-bbcode"},
		},
		func() chroma.Rules {
			return chroma.Rules{
				"root": {},
			}
		},
	)
}

// Name returns the name of the lexer.
func (BBCode) Name() string {
	return heartbeat.LanguageBBCode.StringChroma()
}
