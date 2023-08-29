package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="28" height="30" x="10" y="12" fill="#2F88FF" rx="3"/><path stroke-linecap="round" d="M30 18V10C30 6.68629 27.3137 4 24 4V4C20.6863 4 18 6.68629 18 10V18"/></g>`),
		g.Group(children),
	)
}