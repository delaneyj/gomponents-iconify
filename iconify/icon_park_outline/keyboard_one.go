package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="24" x="4" y="18" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="2"/><circle cx="14" cy="24" r="2" fill="currentColor"/><circle cx="16" cy="30" r="2" fill="currentColor"/><circle cx="10" cy="30" r="2" fill="currentColor"/><circle cx="20" cy="24" r="2" fill="currentColor"/><circle cx="22" cy="30" r="2" fill="currentColor"/><circle cx="26" cy="24" r="2" fill="currentColor"/><circle cx="28" cy="30" r="2" fill="currentColor"/><circle cx="32" cy="24" r="2" fill="currentColor"/><circle cx="34" cy="30" r="2" fill="currentColor"/><circle cx="38" cy="24" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 36h14m2-18v-4.875a1 1 0 0 1 1-1h5a1 1 0 0 0 1-1V6"/></g>`),
		g.Group(children),
	)
}