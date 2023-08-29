package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassEastArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 10v2H5v1H4v1H3v1H2v1H0v-2h1v-1h1v-1h1v-2H2V9H1V8H0V6h2v1h1v1h1v1h1v1m5-4h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}