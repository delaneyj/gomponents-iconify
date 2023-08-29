package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSupplyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M33 34H15L8 17.75L10 9H38L40 17.75L33 34Z"/><path stroke="#000" d="M18 4V9"/><path stroke="#000" d="M30 4V9"/><path stroke="#000" d="M24 34V44H40V36.6316"/><path stroke="#fff" d="M18 21H30"/></g>`),
		g.Group(children),
	)
}