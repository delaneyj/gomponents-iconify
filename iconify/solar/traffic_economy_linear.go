package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficEconomyLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 12A10 10 0 1 1 12 2m3 10h-3m0 0H9m3 0V9m0 3v3m2.5-12.685c3.514.904 6.28 3.67 7.185 7.185"/>`),
		g.Group(children),
	)
}