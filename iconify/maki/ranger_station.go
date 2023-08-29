package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangerStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9 .5l-2 1v3.773L2 8v6h4v-4h3v4h4V8L8 5.273V4l1-.5l2 1l2-1v-3l-2 1l-2-1z"/>`),
		g.Group(children),
	)
}