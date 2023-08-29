package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallSconceFlatVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19v-6H5v6h14m-.27-8.32l1.77-1.77l-1.41-1.41l-1.77 1.77l1.41 1.41m-13.46 0l1.41-1.41L4.91 7.5L3.5 8.91l1.77 1.77M13 8V5h-2v3h2Z"/>`),
		g.Group(children),
	)
}