package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 3a9.5 9.5 0 1 1 0 19a9.5 9.5 0 0 1 0-19Zm0 1a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Z"/>`),
		g.Group(children),
	)
}