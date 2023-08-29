package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 12A10 10 0 1 1 12 2" opacity=".5"/><path d="M14.5 2.315c3.514.904 6.28 3.67 7.185 7.185"/></g>`),
		g.Group(children),
	)
}