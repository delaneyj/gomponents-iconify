package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClarifyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h7v-2H6v2Zm10 0h2V7h-2v10ZM6 13h7v-2H6v2Zm0-4h7V7H6v2ZM2 21V3h20v18H2Zm2-2h16V5H4v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}