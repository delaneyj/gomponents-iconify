package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScoreOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8h16V6H4v2ZM2 20V4h20v8H4v6h4.1v2H2Zm12.95 2l-4.25-4.25l1.4-1.4l2.85 2.8l5.65-5.65l1.4 1.45L14.95 22ZM4 18v-4.5v2.825V6v12Z"/>`),
		g.Group(children),
	)
}