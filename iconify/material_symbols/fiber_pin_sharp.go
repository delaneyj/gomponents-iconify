package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberPinSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h1.5v-2H10V9H5v6Zm6.25 0h1.5V9h-1.5v6ZM14 15h1.25v-3.5L17.8 15H19V9h-1.25v3.5L15.25 9H14v6Zm-7.5-3.5v-1h2v1h-2ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}