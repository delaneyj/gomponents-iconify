package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 3h2v1h1v1h1v1h2v1h2v1h1v1h1v2h-1v2h-1v1h-1v1h-2v1h-1v-1h-1v-1h-1v-2h-1v-1h-1v-1h-1V9H9V8H8V6h1V5h1V4h1m-1 6v1h1v2h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v-1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1"/>`),
		g.Group(children),
	)
}