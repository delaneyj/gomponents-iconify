package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransactionOrder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="30" height="36" x="9" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" d="M18 4V10"/><path stroke="#000" stroke-linecap="round" d="M30 4V10"/><path stroke="#fff" stroke-linecap="round" d="M16 19L32 19"/><path stroke="#fff" stroke-linecap="round" d="M16 27L28 27"/><path stroke="#fff" stroke-linecap="round" d="M16 35H24"/></g>`),
		g.Group(children),
	)
}