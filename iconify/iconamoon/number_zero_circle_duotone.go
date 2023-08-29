package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}