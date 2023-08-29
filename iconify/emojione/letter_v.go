package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M38.4 17.5h6.3l-9.9 29h-5.7l-9.8-29h6.5l6.3 22l6.3-22"/>`),
		g.Group(children),
	)
}