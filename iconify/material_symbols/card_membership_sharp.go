package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMembershipSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v15h-6v5l-4-2l-4 2v-5H2V2Zm2 11h16v-3H4v3Z"/>`),
		g.Group(children),
	)
}