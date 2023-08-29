package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9h4v1H9.71l6.71 6.72l-.7.7L9 10.71V13H8V9M3 4h12v9l-1-1V5H4v10h7l1 1H3V4m17 5v12H8v-3h1v2h10V10h-2V9h3Z"/>`),
		g.Group(children),
	)
}