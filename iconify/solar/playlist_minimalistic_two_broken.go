package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistMinimalisticTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9 14H3m5 4H3m14-1.5V8"/><circle cx="14.5" cy="16.5" r="2.5"/><path stroke-linecap="round" d="M21 12a4 4 0 0 1-4-4M3 6h7.412M15 6h-1.588M13 10H6.824M3 10h1.324"/></g>`),
		g.Group(children),
	)
}