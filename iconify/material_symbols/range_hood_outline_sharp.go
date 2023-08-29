package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangeHoodOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.412-.587Q2 18.825 2 18v-4.2q0-.375.15-.738q.15-.362.45-.662L7 8V3h10v5l4.425 4.425q.275.275.425.637q.15.363.15.763V18q0 .825-.587 1.413Q20.825 20 20 20Zm1.8-8h12.4L15 8.8V5H9v3.8ZM4 18h16v-4H4v4Zm6-1.3v-1.5h4v1.5Z"/>`),
		g.Group(children),
	)
}