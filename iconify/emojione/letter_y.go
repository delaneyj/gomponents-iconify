package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M38 17.5h6.9l-9.7 18.1v10.9h-6.1V35.6l-10-18.1h7.1l6 12.6L38 17.5z"/>`),
		g.Group(children),
	)
}