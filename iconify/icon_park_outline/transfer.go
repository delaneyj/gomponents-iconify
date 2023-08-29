package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 28.392v7.015C37 39.048 34.089 42 30.497 42s-6.503-2.952-6.503-6.593l.012-22.264C24.006 9.198 21.095 6 17.503 6C13.912 6 11 9.198 11 13.143v15.25"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m43 31l-6-6l-6 6m12-20.273C43 15.182 37 19 37 19s-6-3.818-6-8.273c0-1.519.632-2.975 1.757-4.05A6.148 6.148 0 0 1 37 5a6.15 6.15 0 0 1 4.243 1.677A5.598 5.598 0 0 1 43 10.727Zm-26 24C17 39.182 11 43 11 43s-6-3.818-6-8.273c0-1.519.632-2.975 1.757-4.05A6.148 6.148 0 0 1 11 29a6.15 6.15 0 0 1 4.243 1.677A5.598 5.598 0 0 1 17 34.727Z"/><circle cx="37" cy="11" r="2" fill="currentColor"/><circle cx="11" cy="35" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}