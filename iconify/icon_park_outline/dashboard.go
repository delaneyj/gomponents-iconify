package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8.444 41.556A21.931 21.931 0 0 1 2 26C2 13.85 11.85 4 24 4s22 9.85 22 22a21.931 21.931 0 0 1-6.444 15.556"/><path d="M14.1 35.9A13.956 13.956 0 0 1 10 26c0-7.732 6.268-14 14-14"/><path stroke-linejoin="round" d="M24 26v-8"/></g>`),
		g.Group(children),
	)
}