package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 3l1-1l1 1v2h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-7v3a2 2 0 0 1 2 2H9a2 2 0 0 1 2-2v-3H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7V3M6 8v2h12V8H6m0 4v2h7v-2H6Z"/>`),
		g.Group(children),
	)
}