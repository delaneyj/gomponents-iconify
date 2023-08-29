package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReduceDecimalPlaces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path d="M27 9C27 6.23858 24.7614 4 22 4C19.2386 4 17 6.23858 17 9V19C17 21.7614 19.2386 24 22 24C24.7614 24 27 21.7614 27 19V9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M25 32L19 38L25 44"/><path stroke-linecap="round" d="M8 24H9"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 38H19"/></g>`),
		g.Group(children),
	)
}