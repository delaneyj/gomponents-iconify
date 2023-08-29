package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="14" r="3" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="9" r="2" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}