package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.034 11.117A4.002 4.002 0 0 0 6 19h11a5 5 0 1 0-1.17-9.862L14.5 9.5"/><path d="M15.83 9.138a5.5 5.5 0 0 0-10.796 1.98S5.187 12 5.5 12.5"/></g>`),
		g.Group(children),
	)
}