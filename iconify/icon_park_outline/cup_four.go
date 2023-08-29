package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M14 25c0 2 1.071 7 15 7s15-5 15-7V10H14v15Z"/><path stroke-linecap="round" d="M29 16h-6v5l3 3l3-3v-5Zm-3 0v-6M15 40h28"/><path d="M14 14H4s1 5 2 8c.998 3 8 2 8 2"/></g>`),
		g.Group(children),
	)
}