package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaSignalTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 9s1 1.125 1 3s-1 3-1 3m-3-2.99l.01-.011M17 7s2 1.786 2 5s-2 5-2 5M9 9s-1 1.125-1 3s1 3 1 3M7 7s-2 1.786-2 5s2 5 2 5"/></g>`),
		g.Group(children),
	)
}