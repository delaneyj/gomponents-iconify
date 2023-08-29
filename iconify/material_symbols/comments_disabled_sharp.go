package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentsDisabledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-5.3-5.3H2V4.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM22 19.125L16.875 14H18v-2h-3.125l-1-1H18V9h-6.125l-1-1H18V6H8.875l-4-4H22v17.125ZM6 14h5.175l-2-2H6v2Zm0-3h2.175l-2-2H6v2Z"/>`),
		g.Group(children),
	)
}