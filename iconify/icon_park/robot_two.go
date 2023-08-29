package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M5 35C5 33.8954 5.89543 33 7 33H41C42.1046 33 43 33.8954 43 35V42H5V35Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 18L34 18L28 12L34 6L42 6"/><circle cx="8" cy="12" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12L28 12"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 16L18 33"/></g>`),
		g.Group(children),
	)
}