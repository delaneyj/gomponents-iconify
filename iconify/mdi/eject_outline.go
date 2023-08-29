package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17h14v2H5v-2m7-12L5.33 15h13.34L12 5m0 3.6l2.93 4.4H9.07L12 8.6Z"/>`),
		g.Group(children),
	)
}