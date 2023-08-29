package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeekendSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V10h4v6h14v-6h4v10H1Zm6-6V8H4V4h16v4h-3v6H7Z"/>`),
		g.Group(children),
	)
}