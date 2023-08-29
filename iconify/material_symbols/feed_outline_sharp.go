package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h13l5 5v13H3Zm2-2h14V9h-4V5H5v14Zm2-2h10v-2H7v2Zm0-8h5V7H7v2Zm0 4h10v-2H7v2ZM5 5v4v-4v14V5Z"/>`),
		g.Group(children),
	)
}