package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M23 8H43"/><path d="M14 41L6 33"/><path d="M14 7V41"/><path d="M23 18H39"/><path d="M23 28H35"/><path d="M23 38H31"/></g>`),
		g.Group(children),
	)
}