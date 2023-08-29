package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Min(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 4v40h40"/><path d="M10 4s5.313 34 17 34C38.688 38 44 4 44 4"/><path stroke-dasharray="2 6" d="M10 38h34"/></g>`),
		g.Group(children),
	)
}