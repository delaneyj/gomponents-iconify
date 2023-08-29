package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RipplesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.55 5q-.275.45-.413.963T13 7q0 1.65 1.175 2.825T17 11q.525 0 1.038-.138T19 10.45V5h-5.45ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}