package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Find(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M4 7h40M4 23h11M4 39h11"/><path d="M31.5 34a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z"/><path stroke-linecap="round" d="m37 32l7 7.05"/></g>`),
		g.Group(children),
	)
}