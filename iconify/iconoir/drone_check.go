package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="m16 20l2 2l4-4"/></g>`),
		g.Group(children),
	)
}