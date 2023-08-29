package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RateReviewOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h2.45l6.275-6.225l-2.5-2.5L6 11.55V14Zm4.5 0H18v-2h-5.5l-2 2ZM2 22V2h20v16H6l-4 4Zm2-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}