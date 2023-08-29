package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHeaderIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h2v6h4V4h2v14h-2v-6H6v6H4V4m10.59 3.41L18.17 11l-3.58 3.59L16 16l5-5l-5-5l-1.41 1.41Z"/>`),
		g.Group(children),
	)
}