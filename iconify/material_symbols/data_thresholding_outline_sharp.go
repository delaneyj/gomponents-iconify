package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataThresholdingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm14.275-2H19v-1.725L17.275 19ZM5.85 19h1.825l3-3H12.8l-3 3h1.6l3-3h2.125l-3 3h1.625l3-3H19V5H5v12.725L6.725 16H8.85l-3 3Zm1.8-5l-1.4-1.4l4.425-4.425l2 2L16.35 6.5l1.4 1.4l-5.075 5.1l-2-2l-3.025 3ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}