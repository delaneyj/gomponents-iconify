package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M14 21H5V5h16v9"/><path stroke-linecap="round" d="M32 27h11v16H27V32"/><path d="M14 32V14h18v18H14Z"/></g>`),
		g.Group(children),
	)
}