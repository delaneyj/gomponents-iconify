package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquareClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h2v2H4V4m2 16H4v-2h2v2M18 8v8h-2v2H8v-2H6V8h2V2H2v6h2v8H2v6h6v-2h8v2h6v-6h-2V8h2V2h-6v6h2m2 12h-2v-2h2v2M18 6V4h2v2h-2m-4 0h-4V4h4v2Z"/>`),
		g.Group(children),
	)
}