package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FNineKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path d="M30 24a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M26 29.124C26.618 30.81 28.108 32 29.848 32C32.14 32 34 30.549 34 28v-8m-13-4h-7v16m0-8h7"/></g>`),
		g.Group(children),
	)
}