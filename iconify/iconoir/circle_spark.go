package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSpark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2"/><path d="M13 6.5c3.134 0 4.5-1.318 4.5-4.5c0 3.182 1.357 4.5 4.5 4.5c-3.143 0-4.5 1.357-4.5 4.5c0-3.143-1.366-4.5-4.5-4.5Z"/></g>`),
		g.Group(children),
	)
}