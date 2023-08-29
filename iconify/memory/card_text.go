package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8v2H5V8h12M5 12h10v2H5v-2M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`),
		g.Group(children),
	)
}