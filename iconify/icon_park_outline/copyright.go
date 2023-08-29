package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path d="M24 17h-4v7h4c3 0 4-2 4-3.5S27 17 24 17Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 31v-7m0 0v-7h4c3 0 4 2 4 3.5S27 24 24 24h-1m-3 0h3m5 7l-5-7"/></g>`),
		g.Group(children),
	)
}