package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 9l-1.5 5l-.5 2l-.5-2L24 9h-2v14h2v-8l-.2-2l.6 2l1.6 4.6l1.6-4.6l.6-2l-.2 2v8h2V9zM14 21V9h-2v14h8v-2zm-6 2H2v-2h6v-4H4c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h6v2H4v4h4c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2z"/>`),
		g.Group(children),
	)
}