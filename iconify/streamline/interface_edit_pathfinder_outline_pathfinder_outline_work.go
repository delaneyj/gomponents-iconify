package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPathfinderOutlinePathfinderOutlineWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 9.5a1 1 0 0 0 1-1m-5 1H6m3.5-5V6m-4-1.5a1 1 0 0 0-1 1"/><path d="M.5 1.5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v3h3a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-3h-3a1 1 0 0 1-1-1Zm7.5 3h1.5M4.5 8v1.5"/></g>`),
		g.Group(children),
	)
}