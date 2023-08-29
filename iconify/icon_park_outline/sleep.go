package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M16.866 7.47A17.986 17.986 0 0 0 16 13c0 9.941 8.059 18 18 18a17.94 17.94 0 0 0 7.134-1.47C38.801 36.767 32.012 42 24 42c-9.941 0-18-8.059-18-18c0-7.407 4.473-13.768 10.866-16.53Z"/><path stroke-linecap="round" d="M31.66 10H41l-10 8h10"/></g>`),
		g.Group(children),
	)
}