package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21H7v-5H2v-2h20v2h-5v5m5-16V3H2v2h3v5h14V5h3Z"/>`),
		g.Group(children),
	)
}