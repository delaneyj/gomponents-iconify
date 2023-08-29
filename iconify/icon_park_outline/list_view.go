package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="6" rx="3"/><path d="M4 14h40M20 24h16m-16 8h16m-24-8h2m-2 8h2"/></g>`),
		g.Group(children),
	)
}