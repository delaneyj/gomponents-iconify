package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M10 10C10.5 13 11 14.5 11.5 18C11.9 20.8 12 25.1667 12 27C9.83333 28 5 30 5 34C5 38 10 43 24 43C38 43 43 38 43 34C43 30 36 27 36 27C36 27 36 21.5 36.5 18C37 14.5 37.5 13 38 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 27C36 31 35 35 23.5 35"/><ellipse cx="24" cy="10" rx="14" ry="5"/></g>`),
		g.Group(children),
	)
}