package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixKOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15h1.5v-2.25L16.25 15h1.825l-2.325-3l2.325-3H16.25l-1.75 2.25V9H13v6Zm-6.5 0H11v-3.5H8v-1h3V9H6.5v6ZM8 14v-1.5h1.5V14H8Zm-5 7V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}