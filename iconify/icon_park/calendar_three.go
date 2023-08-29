package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" d="M4 20H44"/><path stroke="#fff" d="M4 32H44"/><path stroke="#000" d="M17 4V12"/><path stroke="#000" d="M31 4V12"/><path stroke="#fff" d="M17 20V44"/><path stroke="#fff" d="M31 20V44"/><path stroke="#000" d="M44 13V39"/><path stroke="#000" d="M4 13L4 39"/><path stroke="#000" d="M14 44H34"/></g>`),
		g.Group(children),
	)
}