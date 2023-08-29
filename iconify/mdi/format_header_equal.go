package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHeaderEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h2v6h4V4h2v14h-2v-6H6v6H4V4m10 6V8h7v2h-7m0 2h7v2h-7v-2Z"/>`),
		g.Group(children),
	)
}