package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorSlidingOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11v2H4v-2h2m16-6h-5v14h5V5M7 5H2v14h5V5m15-2c1.11 0 2 .89 2 2v16H0V5c0-1.11.894-2 2-2h7v16h6V3h7m-2 8h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}