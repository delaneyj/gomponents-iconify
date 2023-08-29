package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterHh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 9v6H9V9H7v14h2v-6h4v6h2V9h-2zm10 4h-4V9h-2v14h2v-8h4v8h2v-8a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}