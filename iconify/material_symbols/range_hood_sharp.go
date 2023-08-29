package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangeHoodSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.05 12L7 8V3h10v5l4 4ZM4 20q-.825 0-1.412-.587Q2 18.825 2 18v-4h20v4q0 .825-.587 1.413Q20.825 20 20 20Zm6-3.3h4v-1.5h-4Z"/>`),
		g.Group(children),
	)
}