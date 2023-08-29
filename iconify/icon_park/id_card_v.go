package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" rx="2"/><path fill="#2F88FF" d="M24 19C26.2091 19 28 17.2091 28 15C28 12.7909 26.2091 11 24 11C21.7909 11 20 12.7909 20 15C20 17.2091 21.7909 19 24 19Z"/><path stroke-linecap="round" d="M30 25C30 21.6863 27.3137 19 24 19C20.6863 19 18 21.6863 18 25"/><path stroke-linecap="round" d="M18 31H30"/><path stroke-linecap="round" d="M18 37H25"/></g>`),
		g.Group(children),
	)
}