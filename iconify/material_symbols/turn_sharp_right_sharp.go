package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSharpRightSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-8h10V6.8l-1.6 1.6L13 7l4-4l4 4l-1.4 1.4L18 6.8V15H8v6H6Z"/>`),
		g.Group(children),
	)
}