package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 4v40h40"/><path d="M10 38S15.313 4 27 4c11.688 0 17 34 17 34"/><path stroke-dasharray="2 6" d="M10 4h34"/></g>`),
		g.Group(children),
	)
}