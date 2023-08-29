package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="34" x="8" y="5" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 39v4m20-4v4"/><circle cx="34" cy="33" r="2" fill="currentColor"/><circle cx="14" cy="33" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 23h32M8 21v4m32-4v4M18 13h12"/></g>`),
		g.Group(children),
	)
}