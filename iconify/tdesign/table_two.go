package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 3.5h21v17h-21v-17Zm2 2v3H11v-3H3.5Zm9.5 0v3h7.5v-3H13Zm7.5 5H13v3h7.5v-3Zm0 5H13v3h7.5v-3Zm-9.5 3v-3H3.5v3H11Zm-7.5-5H11v-3H3.5v3Z"/>`),
		g.Group(children),
	)
}