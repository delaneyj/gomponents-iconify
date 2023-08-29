package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4a2 2 0 0 0-2 2v4h2V6h8v3h-2.5l3.5 3.5L20.5 9H18V6a2 2 0 0 0-2-2H8m-5 8v2h8v-2H3m0 3v2h8v-2H3m10 0v2h8v-2h-8M3 18v2h8v-2H3m10 0v2h8v-2h-8Z"/>`),
		g.Group(children),
	)
}