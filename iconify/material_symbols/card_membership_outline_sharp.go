package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMembershipOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15h16v-2H4v2Zm4 7v-5H2V2h20v15h-6v5l-4-2l-4 2ZM4 10h16V4H4v6Zm0 5V4v11Z"/>`),
		g.Group(children),
	)
}