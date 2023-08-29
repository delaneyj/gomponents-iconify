package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h8v2H8m0-4v-2h8v2H8m0-4v-2h8v2H8m11-4H5l7-7l7 7Z"/>`),
		g.Group(children),
	)
}