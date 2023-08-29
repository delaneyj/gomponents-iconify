package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeatLegroomNormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 12V3H3v9a5 5 0 0 0 5 5h6v-2H8a3 3 0 0 1-3-3m15.5 6H19v-7a2 2 0 0 0-2-2h-5V3H6v8a3 3 0 0 0 3 3h7v7h4.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}