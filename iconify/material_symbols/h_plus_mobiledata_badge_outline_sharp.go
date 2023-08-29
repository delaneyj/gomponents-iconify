package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HPlusMobiledataBadgeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 17h2v-4h4v4h2V7h-2v4h-4V7h-2v10Zm11-2h2v-2h2v-2h-2V9h-2v2h-2v2h2v2ZM1 21V3h22v18H1Zm2-2h18V5H3v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}