package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="20" cy="20" r="16" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path d="M44 18v2h-2v-2h2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 20h2v-2h-2v2Zm0 0c0 9.137-5.57 16.973-13.5 20.298M14 35l-4 9h20l-4-9"/><circle cx="20" cy="20" r="4" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 20c0-5.523 4.477-10 10-10"/></g>`),
		g.Group(children),
	)
}