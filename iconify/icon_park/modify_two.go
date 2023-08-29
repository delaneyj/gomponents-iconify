package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModifyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="40" height="30" x="4" y="9" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16 9V39"/><path stroke="#000" stroke-linecap="round" d="M20 9H12"/><path stroke="#000" stroke-linecap="round" d="M20 39H12"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M23 31L37 17"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M25 19L23 17"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M37 31L35 29"/></g>`),
		g.Group(children),
	)
}