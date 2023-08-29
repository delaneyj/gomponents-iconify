package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.996 1.658l10.416 9.259l-1.329 1.495L20 11.449v10.55H4V11.455l-1.094.957l-1.317-1.505l4.41-3.86V3H8v2.254l3.996-3.596ZM6 9.704V20h3v-6h6v6h3V9.67l-5.996-5.33L7.66 8.252L6 9.704ZM13 20v-4h-2v4h2Z"/>`),
		g.Group(children),
	)
}