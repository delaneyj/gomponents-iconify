package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm9-9h1.5V6H12v6Zm3.5 0h1.7l-2-3l2-3h-1.7l-2 3l2 3ZM7 12h4v-1.5H8.5v-.75H11V6H7v1.5h2.5v.75H7V12Zm12 1v-2.5l-6 6l-4-4l-4 4V19l4-4l4 4l6-6Z"/>`),
		g.Group(children),
	)
}