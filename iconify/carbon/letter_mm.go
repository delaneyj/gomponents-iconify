package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterMm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 13h-8v10h2v-8h2v8h2v-8h2v8h2v-8a2 2 0 0 0-2-2zM12 9l-1.52 5l-.48 1.98L9.54 14L8 9H6v14h2v-8l-.16-2l.58 2L10 19.63L11.58 15l.58-2l-.16 2v8h2V9h-2z"/>`),
		g.Group(children),
	)
}