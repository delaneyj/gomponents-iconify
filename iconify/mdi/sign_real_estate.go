package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRealEstate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8H8c-1.1 0-2 .9-2 2v6a2 2 0 0 0 2 2h10c1.11 0 2-.89 2-2v-6a2 2 0 0 0-2-2m-4 8H8v-2h6v2m4-4H8v-2h10v2m4-6H4v16H2V2h2v2h18v2Z"/>`),
		g.Group(children),
	)
}