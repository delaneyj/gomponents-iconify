package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSafeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7"/><path d="M4 41c0-8.837 8.059-16 18-16m6 4.2c0-1.067 7-3.2 7-3.2s7 2.133 7 3.2c0 8.533-7 12.8-7 12.8s-7-4.267-7-12.8Z"/></g>`),
		g.Group(children),
	)
}