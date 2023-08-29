package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="32" height="24" x="8" y="7" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linecap="round" d="M4 7H44"/><path stroke="#000" stroke-linecap="round" d="M15 41L24 31L33 41"/><path stroke="#fff" stroke-linecap="round" d="M16 13H32"/><path stroke="#fff" stroke-linecap="round" d="M16 19H28"/><path stroke="#fff" stroke-linecap="round" d="M16 25H22"/></g>`),
		g.Group(children),
	)
}