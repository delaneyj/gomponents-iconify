package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateCounterclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 11v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1H5V9h1V7h1V6h2V5h4v1h2v1h1v2h1v4h-1v2h-1v1h-2v1H9v-1H6v2h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v3H0Z"/>`),
		g.Group(children),
	)
}