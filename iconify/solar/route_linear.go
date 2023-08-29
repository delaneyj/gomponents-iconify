package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M6.142 6.142C8.904 3.381 10.284 2 12 2c1.716 0 3.096 1.38 5.858 4.142C20.619 8.904 22 10.284 22 12c0 1.716-1.38 3.096-4.142 5.858C15.096 20.619 13.716 22 12 22c-1.716 0-3.096-1.38-5.858-4.142C3.381 15.096 2 13.716 2 12c0-1.716 1.38-3.096 4.142-5.858Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 11.5L13.333 9M16 11.5L13.333 14M16 11.5h-5.333C9.777 11.5 8 12 8 14"/></g>`),
		g.Group(children),
	)
}