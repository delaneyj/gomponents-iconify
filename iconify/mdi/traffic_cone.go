package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficCone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 15l1 4h3v3H3v-3h3l1-4h10m-2-7l1 4H8l1-4h6m-2-7l1 4h-4l1-4h2Z"/>`),
		g.Group(children),
	)
}