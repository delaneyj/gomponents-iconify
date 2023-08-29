package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenKOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15h1.5v-2.25L16.25 15h1.825l-2.325-3l2.325-3H16.25l-1.75 2.25V9H13v6Zm-5.25 0H9.5l1.45-4.7V9H6.5v1.5h2.65L7.75 15ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}