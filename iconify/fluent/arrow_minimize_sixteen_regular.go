package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMinimizeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.854 2.854a.5.5 0 0 0-.708-.708L10 5.293V3.5a.5.5 0 0 0-1 0v2.9a.6.6 0 0 0 .6.6h2.9a.5.5 0 0 0 0-1h-1.793l3.147-3.146ZM6.5 13a.5.5 0 0 1-.5-.5v-1.793l-3.146 3.147a.5.5 0 0 1-.708-.708L5.293 10H3.5a.5.5 0 0 1 0-1h2.9a.6.6 0 0 1 .6.6v2.9a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}