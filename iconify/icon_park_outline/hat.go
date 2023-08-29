package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 16a8 8 0 0 1 8-8h14a8 8 0 0 1 8 8v16H9V16Z"/><rect width="40" height="8" x="4" y="32" rx="2"/><path d="M9 22h8m14 0h8"/></g>`),
		g.Group(children),
	)
}