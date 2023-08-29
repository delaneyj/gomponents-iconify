package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v4h-1v1h-1v1h-1v1h-1v1h4v2H8v-4h1v-1h1v-1h1V9h1V8H8V6Z"/>`),
		g.Group(children),
	)
}