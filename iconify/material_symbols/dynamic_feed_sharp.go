package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFeedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-9h2v7h9v2H2Zm4-4V8h2v7h9v2H6Zm4-4V3h12v10H10Zm2-2h8V7h-8v4Z"/>`),
		g.Group(children),
	)
}