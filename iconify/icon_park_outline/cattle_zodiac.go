package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CattleZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="m38 31l6 6M5 44V18C5 13 7 8 16 6l14-2"/><path stroke-linejoin="round" d="M19 20c1.5 1.5 3.5 5 9 5c3.167 0 10 1.5 10 8v11M16 6c6 0 9 3 9 10"/><path d="M30 44a8 8 0 1 0-16 0"/></g>`),
		g.Group(children),
	)
}