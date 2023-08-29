package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhatThreeWords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 34.91l8.18-21.82m2.73 21.82l8.18-21.82m2.73 21.82L39 13.09"/>`),
		g.Group(children),
	)
}