package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gongfu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="22" cy="8" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 18H18V29H11V43"/><path stroke-linecap="round" stroke-linejoin="round" d="M36.1818 18H26V28.8596L40 43"/></g>`),
		g.Group(children),
	)
}