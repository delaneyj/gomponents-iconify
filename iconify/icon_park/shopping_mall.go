package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingMall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" fill-rule="evenodd" d="M8 44V6C8 5.44772 8.44772 5 9 5H29C29.5523 5 30 5.44772 30 6V44" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 44V6C8 5.44772 8.44772 5 9 5H29C29.5523 5 30 5.44772 30 6V44"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 15L40 20.9993V44"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M4 44H44"/></g>`),
		g.Group(children),
	)
}