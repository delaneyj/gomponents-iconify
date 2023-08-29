package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instruction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" stroke-linejoin="round" rx="2"/><path stroke-linejoin="round" d="M4 10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v6H4v-6Z"/><path d="m25 23l-2 11"/><path stroke-linejoin="round" d="m31 23l6 5l-6 6M17 23l-6 5l6 6"/></g>`),
		g.Group(children),
	)
}