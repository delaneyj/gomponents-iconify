package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedReality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M10.971 6.617a2 2 0 0 1 2.058 0l3.486 2.092a1 1 0 0 1 .485.857v4.302a2 2 0 0 1-.971 1.715l-3 1.8a2 2 0 0 1-2.058 0l-3-1.8A2 2 0 0 1 7 13.868V9.566a1 1 0 0 1 .486-.857l3.485-2.092Z"/><path d="m7 9l5 2.759m0 0L17 9m-5 2.759V17"/><path stroke-linecap="round" d="M6 2H4a2 2 0 0 0-2 2v2m16 16h2a2 2 0 0 0 2-2v-2m0-12V4a2 2 0 0 0-2-2h-2M2 18v2a2 2 0 0 0 2 2h2"/></g>`),
		g.Group(children),
	)
}