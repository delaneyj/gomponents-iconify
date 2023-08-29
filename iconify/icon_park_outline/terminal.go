package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" rx="2"/><path stroke-linecap="round" d="m12 18l7 6l-7 6m11 2h13"/></g>`),
		g.Group(children),
	)
}