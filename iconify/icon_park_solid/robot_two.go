package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRobotTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="M5 35a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v7H5v-7Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 18h-8l-6-6l6-6h8"/><circle cx="8" cy="12" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12h16m-18 4l8 17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRobotTwo0)"/>`),
		g.Group(children),
	)
}