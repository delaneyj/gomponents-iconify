package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewsmodeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h20v18H2Zm2-2h16V5H4v14Zm2-2h12v-2H6v2Zm0-4h4V7H6v6Zm6 0h6v-2h-6v2Zm0-4h6V7h-6v2ZM4 19V5v14Z"/>`),
		g.Group(children),
	)
}