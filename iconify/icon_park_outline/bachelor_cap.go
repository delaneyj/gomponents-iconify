package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BachelorCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m4 13l20-5l20 5l-20 5l-20-5Z"/><path d="M13 16v9.97S18 29 24 29s11-3.03 11-3.03V16M7 14v22m-3-2h6v6H4z"/></g>`),
		g.Group(children),
	)
}