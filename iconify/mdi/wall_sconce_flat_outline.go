package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallSconceFlatOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5v6h14V5H5m12 4H7V7h10v2M5.27 13.32L3.5 15.09l1.41 1.41l1.77-1.77l-1.41-1.41m13.46 0l-1.41 1.41l1.77 1.77l1.41-1.41l-1.77-1.77M11 16v3h2v-3h-2Z"/>`),
		g.Group(children),
	)
}