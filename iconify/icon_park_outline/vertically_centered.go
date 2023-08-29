package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticallyCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M6 7h36"/><path stroke-linejoin="round" d="M16 16h16v16H16z"/><path d="M6 41h36"/></g>`),
		g.Group(children),
	)
}