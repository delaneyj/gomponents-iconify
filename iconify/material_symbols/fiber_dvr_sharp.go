package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberDvrSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15h4.25l.75-.75v-4.5L8.25 9H4v6Zm1.5-1.5v-3h2v3h-2Zm5.6 1.5h1.5l1.75-6h-1.5l-1 3.45l-1-3.45h-1.5l1.75 6Zm3.9 0h1.5v-2h1.15l.85 2H20l-.9-2.1h.9V9h-5v6Zm1.5-3.5v-1h2v1h-2ZM1 20V4h22v16H1Z"/>`),
		g.Group(children),
	)
}