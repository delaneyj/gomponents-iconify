package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3v2H8V3h8m0 4v2H8V7h8m0 4v2H8v-2h8M5 15h14l-7 7l-7-7Z"/>`),
		g.Group(children),
	)
}