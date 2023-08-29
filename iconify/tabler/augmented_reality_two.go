package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedRealityTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 21H8a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v3.5"/><path d="m17 17l-4-2.5l4-2.5l4 2.5V19l-4 2.5z"/><path d="M13 14.5V19l4 2.5m0-4.5l4-2.5M11 4h2"/></g>`),
		g.Group(children),
	)
}