package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrdigeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M8 12v26m32-26v26"/><path d="M8 38s8-11 16-11s16 11 16 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 27h40M4 19h40"/><path stroke-linecap="round" d="M24 15v12m-8-8v8m16-8v8"/></g>`),
		g.Group(children),
	)
}