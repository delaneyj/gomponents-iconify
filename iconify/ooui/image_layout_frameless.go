package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageLayoutFrameless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 3H1v14h18zM3 14l3.5-4.5l2.5 3L12.5 8l4.5 6z"/><path fill="currentColor" d="M19 5H1V3h18zm0 12H1v-2h18z"/>`),
		g.Group(children),
	)
}