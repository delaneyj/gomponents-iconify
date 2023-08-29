package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterBAndWSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3V3h18v18ZM5 19h7v-8l7 8V5h-7v6l-7 8Z"/>`),
		g.Group(children),
	)
}