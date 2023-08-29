package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointOnMapPerspectiveLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21.121 21.121C22 20.243 22 18.828 22 16c0-2.828 0-4.243-.879-5.121m0 10.242C20.243 22 18.828 22 16 22H8c-2.828 0-4.243 0-5.121-.879m18.242 0Zm0-10.242C20.243 10 18.828 10 16 10H8c-2.828 0-4.243 0-5.121.879m18.242 0Zm-18.242 0C2 11.757 2 13.172 2 16c0 2.828 0 4.243.879 5.121m0-10.242Zm0 10.242Z"/><path stroke-linecap="round" d="M21 21L3 11m.5 10l8.5-5"/><circle cx="17" cy="5" r="3"/><path stroke-linecap="round" d="M17 13V8"/></g>`),
		g.Group(children),
	)
}