package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44h40"/><path d="m4 26l8 2v10H4V26Zm16-2l8-4v18h-8V24Zm16-8l8-4v26h-8V16Z"/><path stroke-linecap="round" d="m4 18l8 2L44 4H34"/></g>`),
		g.Group(children),
	)
}