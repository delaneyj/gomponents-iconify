package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M23 9H43"/><path d="M5 16L13 8"/><path d="M13 8V42"/><path d="M23 19H39"/><path d="M23 29H35"/><path d="M23 39H31"/></g>`),
		g.Group(children),
	)
}