package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsAltOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-4 4L1.2 12l4-4L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM12 17l1.1-1.1l-1.425-1.4l-1.075 1.1L12 17Zm6.85-1.05L22.8 12L12 1.2L8.05 5.15l2.9 2.9L12 7l5 5l-1.05 1.05l2.9 2.9ZM7 13h3.175l-2-2H7v2Z"/>`),
		g.Group(children),
	)
}