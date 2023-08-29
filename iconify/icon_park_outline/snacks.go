package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snacks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 14h36V8h-4l-2-4H12l-2 4H6v6Z"/><path stroke-linecap="round" d="m36 44l2-30H10l2 30h24Z"/></g>`),
		g.Group(children),
	)
}