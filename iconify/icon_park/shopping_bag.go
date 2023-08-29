package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 17H43L38.8 43H9.2L5 17Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 17C35 10.3726 30.0751 5 24 5C17.9249 5 13 10.3726 13 17"/><circle cx="17" cy="26" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 33C18 33 20 36 24 36C28 36 30 33 30 33"/><circle cx="31" cy="26" r="2" fill="#fff"/></g>`),
		g.Group(children),
	)
}