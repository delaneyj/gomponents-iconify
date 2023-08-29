package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPathfinderIntersectPathfinderIntersectWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5a1 1 0 0 0 1 1m0-9a1 1 0 0 0-1 1m9 0a1 1 0 0 0-1-1M4 .5h2M.5 4v2m4 6.5a1 1 0 0 0 1 1m8-8a1 1 0 0 0-1-1m0 9a1 1 0 0 0 1-1m-5.5 1h2M13.5 8v2m-4-1.5v-4h-4a1 1 0 0 0-1 1v4h4a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}