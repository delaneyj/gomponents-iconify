package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="34" x="8" y="5" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 39L14 43"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 39L34 43"/><circle cx="34" cy="33" r="2" fill="#fff"/><circle cx="14" cy="33" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 23H40"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 21L8 25"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 21L40 25"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 13H30"/></g>`),
		g.Group(children),
	)
}