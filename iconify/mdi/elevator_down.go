package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 10l4-4H8V2H6v4H3l4 4m10 0l-4-4h3V2h2v4h3l-4 4M7 12h10a2 2 0 0 1 2 2v6c0 1.11-.89 2-2 2H7a2 2 0 0 1-2-2v-6c0-1.1.9-2 2-2m0 2v6h10v-6H7Z"/>`),
		g.Group(children),
	)
}