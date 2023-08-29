package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1Z"/>`),
		g.Group(children),
	)
}