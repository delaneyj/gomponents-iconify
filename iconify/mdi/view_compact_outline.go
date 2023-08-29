package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompactOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v14h19V5H3m2 2h15v4H5V7m0 10v-4h4v4H5m6 0v-4h9v4h-9Z"/>`),
		g.Group(children),
	)
}