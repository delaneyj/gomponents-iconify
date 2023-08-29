package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewTimelineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm3-4h6v-2H6v2Zm6-8h6V7h-6v2Zm-3 4h6v-2H9v2Z"/>`),
		g.Group(children),
	)
}