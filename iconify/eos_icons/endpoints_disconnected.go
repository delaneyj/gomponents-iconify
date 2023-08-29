package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndpointsDisconnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.39 4.93l6.08 6.08H3.01v2H7V18H6v4h3.99v-4h-1v-4.99h1.48l8.65 8.65l1.28-1.28L3.67 3.66L2.39 4.93zm14.62 6.08v-5h1v-4h-4v4h1v5h-1.44l2 2h5.44v-2h-4z"/>`),
		g.Group(children),
	)
}