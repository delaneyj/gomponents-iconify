package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayerStopDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><rect width="14" height="14" x="5" y="5" fill="currentColor" opacity=".16" rx="2"/><rect width="14" height="14" x="5" y="5" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/></g>`),
		g.Group(children),
	)
}