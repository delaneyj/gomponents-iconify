package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2L20 4Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2L20 4Z"/></g>`),
		g.Group(children),
	)
}