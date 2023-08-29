package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHeaderDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h2v6h4V4h2v14h-2v-6H6v6H4V4m16.42 3.41L16.83 11l3.59 3.59L19 16l-5-5l5-5l1.42 1.41Z"/>`),
		g.Group(children),
	)
}