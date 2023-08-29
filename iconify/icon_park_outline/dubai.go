package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dubai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M14 4v40"/><path d="M14.5 6S28 13 32 22s1 22 1 22"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40"/><path stroke-linecap="round" d="M10 15h22m-18 7h8m-8 7h12m-12 7h13"/></g>`),
		g.Group(children),
	)
}