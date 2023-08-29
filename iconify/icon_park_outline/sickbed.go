package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sickbed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m4 23l36 12"/><circle cx="12" cy="16" r="3"/><path stroke-linejoin="round" d="M29 36v-4.5L19 28v8h10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 27.5L23 14l18.374 7a3 3 0 0 1 1.8 3.686L40 35"/></g>`),
		g.Group(children),
	)
}