package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterVariantMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 8H3V6h18v2m-7.19 8H10v2h3.09c.12-.72.37-1.39.72-2M18 11H6v2h12v-2m5 7h-8v2h8v-2Z"/>`),
		g.Group(children),
	)
}