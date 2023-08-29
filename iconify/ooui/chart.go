package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3H1v16h18v-2H3z"/><path fill="currentColor" d="M11 11L8 9l-4 4v3h14V5z"/>`),
		g.Group(children),
	)
}