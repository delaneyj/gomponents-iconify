package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" rx="2"/><path d="M4 14h40m0-3v12m-32-1h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4M4 11v12"/></g>`),
		g.Group(children),
	)
}