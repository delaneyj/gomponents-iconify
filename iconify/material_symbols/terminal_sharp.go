package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h16V8H4v10Zm3.5-1l-1.4-1.4L8.675 13l-2.6-2.6L7.5 9l4 4l-4 4Zm4.5 0v-2h6v2h-6Z"/>`),
		g.Group(children),
	)
}