package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRoll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 6.9V2h-2V0H1v1H0v3h1v1h13V3h1v3.1l-8 1V9H6v.9s.5 0 .5.9s-.5.6-.5 1.5v2.8s0 .9 1.5.9s1.5-.9 1.5-.9v-2.8c0-.9-.5-.7-.5-1.5s.5-.9.5-.9V9H8V7.9l8-1z"/>`),
		g.Group(children),
	)
}