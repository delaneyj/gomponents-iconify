package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeRhombus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.586 12.5L11.5 3.586l8.914 8.914l-8.914 8.914L2.586 12.5ZM11.5 5L4 12.5l7.5 7.5l7.5-7.5L11.5 5Z"/>`),
		g.Group(children),
	)
}