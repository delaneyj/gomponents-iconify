package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snacks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M6 14H42V8H38L36 4H12L10 8H6V14Z"/><path stroke-linecap="round" d="M36 44L38 14H10L12 44H36Z"/></g>`),
		g.Group(children),
	)
}