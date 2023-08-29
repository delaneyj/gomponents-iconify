package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="26" height="40" x="11" y="4" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M22 10L26 10"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M20 38H28"/></g>`),
		g.Group(children),
	)
}