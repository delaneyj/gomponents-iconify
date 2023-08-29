package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTowerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M12 44L20 4H28L36 44"/><path d="M15 10H24H33"/><path stroke-linejoin="round" d="M12 26L36 26"/><path stroke-linejoin="round" d="M15 27L35 39"/><path stroke-linejoin="round" d="M33 27L14 39"/><path stroke-linejoin="round" d="M30 11L15 26"/><path stroke-linejoin="round" d="M18 11L33 26"/></g>`),
		g.Group(children),
	)
}