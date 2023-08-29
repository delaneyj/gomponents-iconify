package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsVerticalClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19V3H4v16H2v2h20v-2h-2M13 5h1.5v14H13V5m-2 14H9.5V5H11v14M6 5h1.5v14H6V5m10.5 14V5H18v14h-1.5Z"/>`),
		g.Group(children),
	)
}