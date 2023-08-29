package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsAltOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-4 4L1.2 12l4-4L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM14.6 17.4l-1.5-1.5L12 17l-1.4-1.4l1.075-1.1l-1.5-1.5H7v-2h1.175L6.6 9.4L4 12l8 8l2.6-2.6Zm4.25-1.45l-1.4-1.4L20 12l-8-8l-2.55 2.55l-1.4-1.4L12 1.2L22.8 12l-3.95 3.95Zm-2.9-2.9L17 12l-5-5l-1.05 1.05l5 5Zm-2.5-2.5ZM10.6 13.4Z"/>`),
		g.Group(children),
	)
}