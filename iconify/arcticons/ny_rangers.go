package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NyRangers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24 42.5s1.922-1.898 3.545-1.932l9.92-.056s2.617.468 3.071-2.459l.078-29.508l1.903-2.97l-7.022-.022l-2.149 2.032l-2.316-2.069l-7.622-.016m.592 37s-1.922-1.898-3.545-1.932l-9.92-.056s-2.617.468-3.071-2.459L7.386 8.546L5.483 5.575l7.022-.022l2.149 2.032l2.316-2.069l6.438-.016M7.52 16.244l32.916-.026"/><path d="M7.43 20.792c.398.13 27.006 19.598 27.006 19.598M16.722 16.247l23.582 17.005"/></g>`),
		g.Group(children),
	)
}