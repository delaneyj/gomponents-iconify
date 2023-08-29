package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16.011V8.98L7.98 4h7.04L20 8.98v7.046L15.025 21H7.99L3 16.011ZM8.393 5L4 9.393v6.204L8.403 20h6.208L19 15.611V9.393L14.607 5H8.393Z"/>`),
		g.Group(children),
	)
}