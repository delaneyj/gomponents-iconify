package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSearchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 41C4 32.1634 12.0589 25 22 25"/><circle cx="35" cy="34" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 38L44 41"/></g>`),
		g.Group(children),
	)
}