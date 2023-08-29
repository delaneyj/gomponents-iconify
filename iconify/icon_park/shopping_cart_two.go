package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M5 7H11L18 24H35.5L43 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 12L33 12"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 6V18"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 24L14 30H40"/><circle cx="19" cy="39" r="3" fill="#2F88FF"/><circle cx="35" cy="39" r="3" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}