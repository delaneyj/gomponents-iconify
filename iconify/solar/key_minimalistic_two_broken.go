package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M21.064 12.5A7 7 0 1 0 18 15.326"/><circle cx="15" cy="9" r="2"/><path stroke-linecap="round" d="m3.5 20.5l6-6M6 21l-1.5-1.5m2-2L8 19"/></g>`),
		g.Group(children),
	)
}