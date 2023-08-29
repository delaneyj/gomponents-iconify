package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSmoking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 15v3m0-7c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2m11 9c0-6-4-6-4-6s4 1 4-3m0 13v3"/><path d="M2.6 18H18l-3-3H2.6a.6.6 0 0 0-.6.6v1.8a.6.6 0 0 0 .6.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m3 3l18 18"/></g>`),
		g.Group(children),
	)
}