package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24.778 7c-11.046 0-20 8.954-20 20c0 5.23 1.713 10.436 5 14h30c3.286-3.564 5-8.77 5-14c0-11.046-8.954-20-20-20Z"/><circle cx="24.778" cy="30" r="4"/><path d="M24.778 20v6m0-14v2m-15 14h2m2-10l1.414 1.414M37.778 28h2m-5-8.586L36.192 18"/></g>`),
		g.Group(children),
	)
}