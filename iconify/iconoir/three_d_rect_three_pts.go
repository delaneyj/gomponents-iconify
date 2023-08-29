package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDRectThreePts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 21V3.6a.6.6 0 0 1 .6-.6H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 21h3.4a.6.6 0 0 0 .6-.6V17m0-10v2m0 3v2M7 21h2m3 0h2"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM21 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}