package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 3h10v7H6V3zm9 11a2 2 0 1 1-3.999.001A2 2 0 0 1 15 14zm-2-3c1.3 0 2.4.8 2.8 2h.2v-2h-3z"/><path fill="currentColor" d="M5 5H1L0 9v4h1.2c.4-1.2 1.5-2 2.8-2s2.4.8 2.8 2h3.4c.4-1.2 1.5-2 2.8-2H5V5zM4 9H1l.8-3H4v3z"/><path fill="currentColor" d="M6 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6 14z"/>`),
		g.Group(children),
	)
}