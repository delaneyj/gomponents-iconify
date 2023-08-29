package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 43h38M4 25V14.99c1.216 0 4.818 4.179 6 6.01l11 2V5l4 2l5.831 18.978L40 28c4 1 4 3.5 4 4c0 3-3.5 3-4 3c-4 0-36-10-36-10Z"/>`),
		g.Group(children),
	)
}