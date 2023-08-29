package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPathfinderMergePathfinderMergeWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 9.5h1a1 1 0 0 0 1-1v-1m0-3v-3a1 1 0 0 0-1-1h-7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h3m0 3a1 1 0 0 0 1 1m8-8a1 1 0 0 0-1-1m0 9a1 1 0 0 0 1-1m-4-8h.5m-2 9h2M13.5 8v2m-9-.5v.5"/>`),
		g.Group(children),
	)
}