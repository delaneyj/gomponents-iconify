package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21L11.5 2.813L22 21H1Zm19.268-1L11.5 4.813L2.732 20h17.536Z"/>`),
		g.Group(children),
	)
}