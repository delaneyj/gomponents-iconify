package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 8h40"/><path d="M8 8h32v26H8V8Z" clip-rule="evenodd"/><path d="m31 18l3 3l-3 3m-14 0l-3-3l3-3m-1 24l8-8l8 8"/></g>`),
		g.Group(children),
	)
}