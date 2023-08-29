package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PullDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 8V40L24.2 44V4L6 8Z"/><path stroke="#000" d="M24.2002 8H42.0002V40H24.2002"/><path stroke="#fff" d="M18 22V26"/></g>`),
		g.Group(children),
	)
}