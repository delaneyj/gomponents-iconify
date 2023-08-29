package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12h2L12 2L5.05 12H7l-3.9 6h6.92v4h3.96v-4H21z"/>`),
		g.Group(children),
	)
}