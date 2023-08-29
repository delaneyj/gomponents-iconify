package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignToMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M21 12H3"/><path stroke-linejoin="round" d="M12 2v6m0 14v-6M9 5l3 3l3-3M9 19l3-3l3 3"/></g>`),
		g.Group(children),
	)
}