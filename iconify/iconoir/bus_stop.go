package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m16 16.01l.01-.011M6 16.01l.01-.011M20 22V8m0 0h-2V2h4v6h-2Zm-4 12H2.6a.6.6 0 0 1-.6-.6v-6.8a.6.6 0 0 1 .6-.6H16m-2-4H6m8-6H6a4 4 0 0 0-4 4v2"/><path d="M3.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20m7 0v1.9a.6.6 0 0 0 .6.6h.9"/></g>`),
		g.Group(children),
	)
}