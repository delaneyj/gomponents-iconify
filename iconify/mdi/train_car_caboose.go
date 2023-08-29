package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarCaboose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 9V7h-8V6h1V4H8v2h1v1H1v2h1v6H1v2h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1v-2h-1V9h1M4 15H3V9h1v6m7-3H6V9h5v3m7 0h-5V9h5v3m3 3h-1V9h1v6Z"/>`),
		g.Group(children),
	)
}