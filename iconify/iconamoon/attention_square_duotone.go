package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttentionSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V8"/></g>`),
		g.Group(children),
	)
}