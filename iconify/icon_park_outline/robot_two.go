package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M5 35a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v7H5v-7Zm37-17h-8l-6-6l6-6h8"/><circle cx="8" cy="12" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12h16m-18 4l8 17"/></g>`),
		g.Group(children),
	)
}