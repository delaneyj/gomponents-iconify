package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointOnMapPerspectiveBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 16c0 2.828 0 4.243.879 5.121C3.757 22 5.172 22 8 22h8c2.828 0 4.243 0 5.121-.879C22 20.243 22 18.828 22 16c0-2.828 0-4.243-.879-5.121C20.243 10 18.828 10 16 10H8c-2.828 0-4.243 0-5.121.879c-.3.3-.498.662-.628 1.121M21 21l-3.764-2.091M3 11l11 6.111"/><path d="m3.5 21l6.358-3.74L12 16m5-3V8m-.5-5.959A3 3 0 1 1 14.041 4.5"/></g>`),
		g.Group(children),
	)
}