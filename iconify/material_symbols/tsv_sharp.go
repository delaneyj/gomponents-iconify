package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TsvSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.25 15h1.5v-4.5H9V9H5v1.5h1.25V15Zm3.25 0h4v-3.65H11v-.85h2.5V9h-4v3.6H12v.9H9.5V15Zm6.25 0h1.5L19 9h-1.5l-1 3.45l-1-3.45H14l1.75 6ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}