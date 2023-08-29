package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 10C6 7.79086 7.79086 6 10 6H38C40.2091 6 42 7.79086 42 10V38C42 40.2091 40.2091 42 38 42H10C7.79086 42 6 40.2091 6 38V10Z"/><path stroke="#fff" stroke-linecap="round" d="M6 32H42"/><path stroke="#fff" stroke-linecap="round" d="M20 16L24 20L28 16"/><path stroke="#000" stroke-linecap="round" d="M6 26V38"/><path stroke="#000" stroke-linecap="round" d="M42 26V38"/></g>`),
		g.Group(children),
	)
}