package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceBetweenV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5v4h14V5h-2v2H7V5H5Zm0 14v-4h14v4h-2v-2H7v2H5Zm2-8h10v2H7v-2Z"/>`),
		g.Group(children),
	)
}