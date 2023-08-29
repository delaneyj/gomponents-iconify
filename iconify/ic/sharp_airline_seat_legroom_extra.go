package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAirlineSeatLegroomExtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3H2v14h11v-2H4zm18.24 12.96l-2.53 1.15l-3.41-6.98A2.019 2.019 0 0 0 14.51 9H11V3H5v11h10l3.41 7l5.07-2.32l-1.24-2.72z"/>`),
		g.Group(children),
	)
}