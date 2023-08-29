package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacePowder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><ellipse cx="24" cy="30" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="16" ry="6"/><ellipse cx="24" cy="14" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="16" ry="10"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M23 10L18 13"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 14L25 17"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 38C40 41.3137 32.8366 44 24 44C15.1634 44 8 41.3137 8 38"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 38V30"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 38V30"/><ellipse cx="24" cy="30" fill="#fff" rx="7" ry="2"/></g>`),
		g.Group(children),
	)
}