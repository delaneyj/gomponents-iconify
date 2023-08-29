package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSharpLeftSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-6H6V6.8L4.4 8.4L3 7l4-4l4 4l-1.4 1.4L8 6.8V13h10v8h-2Z"/>`),
		g.Group(children),
	)
}