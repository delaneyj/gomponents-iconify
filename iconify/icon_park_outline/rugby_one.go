package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RugbyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M4 24c2.292 8.63 10.377 15 20 15s17.708-6.37 20-15C41.708 15.37 33.623 9 24 9S6.292 15.37 4 24Z"/><path stroke-linecap="round" d="M40 24h4m-30 0h20M4 24h4m9-4v8m14-8v8m-7-8v8"/></g>`),
		g.Group(children),
	)
}