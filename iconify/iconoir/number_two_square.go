package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 10.8v-.4c0-1.325 1.033-2.4 2.308-2.4c1.274 0 2.307 1.075 2.307 2.4c0 .457-.122.884-.336 1.248C12.73 13.44 9.5 16 9.5 16h5"/></g>`),
		g.Group(children),
	)
}