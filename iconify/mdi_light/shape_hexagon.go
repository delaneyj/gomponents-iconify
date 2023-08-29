package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeHexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.593 21l-4.905-8.494L6.6 4h9.808l4.908 8.5l-4.908 8.5H6.593ZM15.83 5H7.177l-4.334 7.506L7.17 20h8.66l4.33-7.5L15.83 5Z"/>`),
		g.Group(children),
	)
}