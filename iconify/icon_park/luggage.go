package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="32" height="26" x="8" y="14" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="2"/><path stroke="#fff" d="M20 23L20 31"/><path stroke="#000" stroke-linejoin="round" d="M15 40V44"/><path stroke="#000" stroke-linejoin="round" d="M33 40V44"/><path stroke="#fff" d="M28 23V31"/><path stroke="#000" stroke-linejoin="round" d="M19 4H29"/><path stroke="#000" stroke-linejoin="round" d="M24 4L24 14"/></g>`),
		g.Group(children),
	)
}