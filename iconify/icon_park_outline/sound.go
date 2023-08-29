package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><circle cx="24" cy="15" r="5" stroke="currentColor" stroke-width="4"/><circle cx="14" cy="10" r="2" fill="currentColor"/><circle cx="14" cy="38" r="2" fill="currentColor"/><circle cx="34" cy="10" r="2" fill="currentColor"/><circle cx="34" cy="38" r="2" fill="currentColor"/><circle cx="24" cy="32" r="6" stroke="currentColor" stroke-width="4"/></g>`),
		g.Group(children),
	)
}