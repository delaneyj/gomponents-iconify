package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPathfinderDividePathfinderDivideWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5a1 1 0 0 0 1 1m0-9a1 1 0 0 0-1 1m9 0a1 1 0 0 0-1-1M4 .5h2M.5 4v2m9-1.5h3a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-3h4a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}