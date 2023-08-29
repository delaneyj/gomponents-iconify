package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 16c0-6.075-3.582-12-8-12s-8 5.925-8 12m-7 0h30l1 12H27v-3h-6v3H8l1-12ZM8 28L6 42h36l-2-14"/><path d="M21 25h6v6h-6v-6Z"/></g>`),
		g.Group(children),
	)
}