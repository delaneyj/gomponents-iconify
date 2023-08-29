package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicyclingLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="15" cy="4" r="2"/><circle cx="6" cy="18" r="3" opacity=".4"/><circle cx="18" cy="18" r="3" opacity=".4"/><path stroke-linecap="round" d="M18.5 10h-2.026c-1.22 0-1.831 0-2.38-.227c-.55-.227-.983-.657-1.85-1.518l-.576-.574c-.785-.78-1.178-1.17-1.642-1.127c-.465.042-.78.497-1.411 1.406L7.388 9.727c-.642.925-.963 1.388-.835 1.843c.13.455.646.68 1.678 1.13l1.476.644c1.364.595 2.046.893 2.374 1.493c.328.6.21 1.335-.024 2.804L12 18"/></g>`),
		g.Group(children),
	)
}